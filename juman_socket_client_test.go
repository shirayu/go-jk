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
	conns := s.SocketClientConns(server)
	for {
		go s.HandleConn(<-conns)
	}
	return nil
}

func (s *JumanTestServer) SocketClientConns(listener net.Listener) chan net.Conn {
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
		} else if string(line) == input_sample+"\n" {
			fmt.Fprintf(client, juman_input_sample+"\n")
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

	juman, err := NewJumanSocketClient("localhost:" + strconv.Itoa(server.Port()))
	if err != nil {
		t.Fatal("Error to open the juman socket: ", err)
	}

	ret_lines, err := juman.RawParse(input_sample)
	if err != nil {
		t.Error("Error to parse [%v]", err)
	}
	if c := strings.Count(ret_lines, "\n"); c != 4 {
		t.Errorf("expceted length is 4 but %d", c)
	}

	s, err := juman.Parse(input_sample)
	if err != nil {
		t.Error("Error to parse [%v]", err)
	}
	if s.Len() != 4 {
		t.Errorf("expceted length is 4 but %d", s.Len())
	}

	fin <- 1
	server.Wait()
}
