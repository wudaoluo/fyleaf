package console

import (
	"bufio"
	"github.com/wudaoluo/fyleaf/network"
	"github.com/wudaoluo/fyleaf/conf"
	"strings"
)


type Agent struct {
	conn   *network.TCPConn
	reader *bufio.Reader
}


func newAgent(conn *network.TCPConn) network.Agent {
	a := new(Agent)
	a.conn = conn
	a.reader = bufio.NewReader(conn)
	return a
}


func (a *Agent) Run() {
	for {
		a.conn.Write([]byte(conf.ConsolePrompt))


		//按照 \n 分割,阻塞直到收集到 \n 字符
		line, err := a.reader.ReadString('\n')
		if err != nil {
			break
		}

		//去除 \r\n
		line = strings.TrimSuffix(line[:len(line)-1], "\r")


		//字符串按照空格分割
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		if args[0] == "exit" {
			break
		}
		var c Command
		for _, _c := range commands {
			if _c.name() == args[0] {
				c = _c
				break
			}
		}
		if c == nil {
			a.conn.Write([]byte("command not found, try `help` for help\r\n"))
			continue
		}

		output := c.run(args[1:])
		if output != "" {
			a.conn.Write([]byte(output + "\r\n"))
		}
	}
}

func (a *Agent) OnClose() {}
