package jk

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"net/textproto"
	"strings"
)

type SocketClient struct {
	address    string
	connection net.Conn
	tp         *textproto.Reader
}

func NewSocketClient(address string, option string) (*SocketClient, error) {
	var err error
	self := new(SocketClient)
	self.address = address

	self.connection, err = net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	//     defer self.connection.Close()

	r := bufio.NewReader(self.connection)
	self.tp = textproto.NewReader(r)

	fmt.Fprintf(self.connection, option)
	for {
		line, _ := self.tp.ReadLine()
		if strings.Index(line, "OK") != -1 {
			break
		}
	}
	return self, err
}

func (self *SocketClient) RawParse(query string) (string, error) {
	fmt.Fprintf(self.connection, "%s\n", query)

	var buf bytes.Buffer
	for {
		line, err := self.tp.ReadLine()
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
