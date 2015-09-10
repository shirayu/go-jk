package jk

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/lestrrat/go-tcptest"
	"io"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	juman_input_sample = `パン ぱん パン 名詞 6 普通名詞 1 * 0 * 0 "代表表記:パン/ぱん カテゴリ:人工物-食べ物 ドメイン:料理・食事"
が が が 助詞 9 格助詞 1 * 0 * 0 NIL
食べ たべ 食べる 動詞 2 * 0 母音動詞 1 未然形 3 "代表表記:食べる/たべる ドメイン:料理・食事"
られる られる られる 接尾辞 14 動詞性接尾辞 7 母音動詞 1 基本形 2 "代表表記:られる/られる"
EOS`
	knp_output_sample = `# S-ID:1 KNP:4.2-0b760fd DATE:2015/01/01 SCORE:-8.99590
* 1D <BGH:パン/ぱん><文頭><ガ><助詞><体言><係:ガ格><区切:0-0><格要素><連用要素><正規化代表表記:パン/ぱん><主辞代表表記:パン/ぱん>
+ 1D <BGH:パン/ぱん><文頭><ガ><助詞><体言><係:ガ格><区切:0-0><格要素><連用要素><名詞項候補><先行詞候補><正規化代表表記:パン/ぱん><解析格:ガ>
パン ぱん パン 名詞 6 普通名詞 1 * 0 * 0 "代表表記:パン/ぱん カテゴリ:人工物-食べ物 ドメイン:料理・食事" <代表表記:パン/ぱん><カテゴリ:人工物-食べ物><ドメイン:料理・食事><正規化代表表記:パン/ぱん><記英数カ><カタカナ><名詞相当語><文頭><自立><内容語><タグ単位始><文節始><固有キー><文節主辞>
が が が 助詞 9 格助詞 1 * 0 * 0 NIL <かな漢字><ひらがな><付属>
* -1D <BGH:食べる/たべる><文末><態:受動|可能><〜られる><用言:動><レベル:C><区切:5-5><ID:（文末）><提題受:30><主節><動態述語><正規化代表表記:食べる/たべる><主辞代表表記:食べる/たべる>
+ -1D <BGH:食べる/たべる><文末><態:受動|可能><〜られる><用言:動><レベル:C><区切:5-5><ID:（文末）><提題受:30><主節><動態述語><正規化代表表記:食べる/たべる><用言代表表記:食べる/たべる+られる/られる><時制-未来><主題格:一人称優位><格関係0:ガ:パン><格解析結果:食べる/たべる+られる/られる:動1:ガ/C/パン/0/0/1;ニ/U/-/-/-/-;デ/U/-/-/-/-;カラ/U/-/-/-/-;時間/U/-/-/-/-;ノ/U/-/-/-/-;ガ２/U/-/-/-/->
食べ たべ 食べる 動詞 2 * 0 母音動詞 1 未然形 3 "代表表記:食べる/たべる ドメイン:料理・食事" <代表表記:食べる/たべる><ドメイン:料理・食事><正規化代表表記:食べる/たべる><かな漢字><活用語><自立><内容語><タグ単位始><文節始><文節主辞>
られる られる られる 接尾辞 14 動詞性接尾辞 7 母音動詞 1 基本形 2 "代表表記:られる/られる" <代表表記:られる/られる><正規化代表表記:られる/られる><かな漢字><ひらがな><活用語><文末><表現文末><付属>
EOS
`
)

type KnpTestServer struct {
	Host string
	Port int
}

func (s *KnpTestServer) Run() error {
	server, err := net.Listen("tcp", s.Host+":"+strconv.Itoa(s.Port))
	if err != nil {
		return err
	}
	conns := s.ClientConns(server)
	for {
		go s.HandleConn(<-conns)
	}
	return nil
}

func (s *KnpTestServer) ClientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, err := listener.Accept()
			if err != nil {
				fmt.Printf("couldn't accept: " + err.Error())
				continue
			}
			i++
			//             fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func (s *KnpTestServer) HandleConn(client net.Conn) {
	b := bufio.NewReader(client)
	var buf bytes.Buffer
	for {
		line, err := b.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		if strings.HasPrefix(string(line), "RUN") {
			fmt.Fprintf(client, "OK\n")
		} else {
			buf.Write(line)
		}

		if string(line) == "EOS\n" {
			fmt.Fprintf(client, knp_output_sample)
			buf.Reset()
		}
	}
}

func TestKnp(t *testing.T) {
	fin := make(chan int)
	jts := func(port int) {
		jtserver := &KnpTestServer{Host: "localhost", Port: port}
		go jtserver.Run()
		<-fin
		//         jtserver.Shutdown()
	}
	server, err := tcptest.Start(jts, 3*time.Second)
	if err != nil {
		t.Error("Failed to start jtserver: %s", err)
	}

	knp, err := NewKnpClient("localhost:" + strconv.Itoa(server.Port()))
	if err != nil {
		t.Fatal("Error to open the knp socket: ", err)
	}

	ret_lines, err := knp.RawParse(juman_input_sample)
	if err != nil {
		t.Error("Error to parse [%v]", err)
	}
	if c := strings.Count(ret_lines, "\n"); c != 9 {
		t.Errorf("expceted length is 9 but %d", c)
	}

	s, err := knp.Parse(juman_input_sample)
	if err != nil {
		t.Error("Error to parse [%v]", err)
	}
	if s.Len() != 4 {
		t.Errorf("expceted length is 4 but %d", s.Len())
	}

	fin <- 1
	server.Wait()
}
