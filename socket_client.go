package jk

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"net/textproto"
	"strings"
)

//SocketClient communicates to a server with a socket
type SocketClient struct {
	address    string
	connection net.Conn
	tp         *textproto.Reader
}

//NewSocketClient creats a SocketClient
func NewSocketClient(address string, option string) (*SocketClient, error) {
	var err error
	scc := new(SocketClient)
	scc.address = address

	scc.connection, err = net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	//     defer scc.connection.Close()

	r := bufio.NewReader(scc.connection)
	scc.tp = textproto.NewReader(r)

	fmt.Fprintf(scc.connection, option)
	for {
		line, _ := scc.tp.ReadLine()
		if strings.Index(line, "OK") != -1 {
			break
		}
	}
	return scc, err
}

//RawParse returns a single raw result which ends with "EOS" for the given sentence
func (scc *SocketClient) RawParse(query string) (string, error) {
	fmt.Fprintf(scc.connection, "%s\n", query)

	var buf bytes.Buffer
	for {
		line, err := scc.tp.ReadLine()
		if err != nil {
			return "", err
		} else if line == "EOS" {
			buf.WriteString("EOS")
			break
		}

		buf.WriteString(line)
		buf.WriteRune('\n')
	}
	return buf.String(), nil
}
