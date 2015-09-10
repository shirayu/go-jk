package jk

import (
	"bufio"
	"fmt"
	"github.com/lestrrat/go-tcptest"
	"io"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"
)

type JumanTestServer struct {
	Host string
	Port int
}

func (s *JumanTestServer) Run() error {
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

func (s *JumanTestServer) ClientConns(listener net.Listener) chan net.Conn {
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

func (s *JumanTestServer) HandleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		if strings.HasPrefix(string(line), "RUN") {
			fmt.Fprintf(client, "OK\n")
		} else if string(line) == "パンが食べられる\n" {
			out := `パン ぱん パン 名詞 6 普通名詞 1 * 0 * 0 "代表表記:パン/ぱん カテゴリ:人工物-食べ物 ドメイン:料理・食事"
が が が 助詞 9 格助詞 1 * 0 * 0 NIL
食べ たべ 食べる 動詞 2 * 0 母音動詞 1 未然形 3 "代表表記:食べる/たべる ドメイン:料理・食事"
られる られる られる 接尾辞 14 動詞性接尾辞 7 母音動詞 1 基本形 2 "代表表記:られる/られる"
EOS
`
			fmt.Fprintf(client, out)
		} else {
			fmt.Fprintf(client, "Un expected input\n")
			fmt.Fprintf(client, "EOS\n")
		}
	}
}

func TestJuman(t *testing.T) {
	fin := make(chan int)
	jts := func(port int) {
		jtserver := &JumanTestServer{Host: "localhost", Port: port}
		go jtserver.Run()
		<-fin
		//         jtserver.Shutdown()
	}
	server, err := tcptest.Start(jts, 3*time.Second)
	if err != nil {
		t.Error("Failed to start jtserver: %s", err)
	}

	juman, err := NewJumanClient("localhost:" + strconv.Itoa(server.Port()))
	if err != nil {
		t.Fatal("Error to open the juman socket: ", err)
	}

	ret_lines, err := juman.RawParse("パンが食べられる")
	if err != nil {
		t.Error("Error to parse [%v]", err)
	}
	if c := strings.Count(ret_lines, "\n"); c != 4 {
		t.Errorf("expceted length is 4 but %d", c)
	}

	s, err := juman.Parse("パンが食べられる")
	if err != nil {
		t.Error("Error to parse [%v]", err)
	}
	if s.Len() != 4 {
		t.Errorf("expceted length is 4 but %d", s.Len())
	}

	fin <- 1
	server.Wait()
}
