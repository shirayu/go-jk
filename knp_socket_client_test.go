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

type KnpTestServer struct {
	Host string
	Port int
}

func (s *KnpTestServer) Run() error {
	server, err := net.Listen("tcp", s.Host+":"+strconv.Itoa(s.Port))
	if err != nil {
		return err
	}
	conns := s.SocketClientConns(server)
	for {
		go s.HandleConn(<-conns)
	}
	//     return nil
}

func (s *KnpTestServer) SocketClientConns(listener net.Listener) chan net.Conn {
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
			fmt.Fprintf(client, knpOutputSample)
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
		t.Errorf("Failed to start jtserver: %s", err)
	}

	knp, err := NewKnpSocketClient("localhost:" + strconv.Itoa(server.Port()))
	if err != nil {
		t.Fatal("Error to open the knp socket: ", err)
	}

	retLines, err := knp.RawParse(jumanInputSample)
	if err != nil {
		t.Errorf("Error to parse [%v]", err)
	}
	if c := strings.Count(retLines, "\n"); c != 9 {
		t.Errorf("expceted length is 9 but %d", c)
	}

	s, err := knp.Parse(jumanInputSample)
	if err != nil {
		t.Errorf("Error to parse [%v]", err)
	}
	if s.Len() != 4 {
		t.Errorf("expceted length is 4 but %d", s.Len())
	}

	fin <- 1
	server.Wait()
}
