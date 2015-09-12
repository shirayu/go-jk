package jk

import (
	"bufio"
	"bytes"
	"io"
	"os/exec"
)

type CommandClient struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
}

func NewCommandClient(command string, options ...string) (*CommandClient, error) {
	cmd := exec.Command(command, options...)
	self := CommandClient{cmd: cmd}
	var err error
	if self.stdin, err = self.cmd.StdinPipe(); err != nil {
		return nil, err
	}
	if self.stdout, err = self.cmd.StdoutPipe(); err != nil {
		return nil, err
	}
	err = cmd.Start()
	return &self, err
}

func (self *CommandClient) RawParse(query string) (string, error) {
	io.WriteString(self.stdin, query)
	io.WriteString(self.stdin, "\n")

	var buf bytes.Buffer
	scanner := bufio.NewScanner(self.stdout)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "EOS" {
			buf.WriteString("EOS")
			break
		}
		buf.WriteString(line)
		buf.WriteRune('\n')
	}
	return buf.String(), scanner.Err()
}
