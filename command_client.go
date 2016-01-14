package jk

import (
	"bufio"
	"bytes"
	"io"
	"os/exec"
	"sync"
)

//CommandClient execute the given command
type CommandClient struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
	mutex  *sync.Mutex
}

//NewCommandClient creates a new CommandClient
func NewCommandClient(command string, options ...string) (*CommandClient, error) {
	cmd := exec.Command(command, options...)
	cmdcl := CommandClient{cmd: cmd}
	var err error
	if cmdcl.stdin, err = cmdcl.cmd.StdinPipe(); err != nil {
		return nil, err
	}
	if cmdcl.stdout, err = cmdcl.cmd.StdoutPipe(); err != nil {
		return nil, err
	}
	err = cmd.Start()
	cmdcl.mutex = new(sync.Mutex)
	return &cmdcl, err
}

//RawParse returns a single raw result which ends with "EOS" for the given sentence
func (cmdcl *CommandClient) RawParse(query string) (string, error) {
	io.WriteString(cmdcl.stdin, query)
	io.WriteString(cmdcl.stdin, "\n")

	var buf bytes.Buffer
	cmdcl.mutex.Lock()
	defer cmdcl.mutex.Unlock()
	scanner := bufio.NewScanner(cmdcl.stdout)
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
