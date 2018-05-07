package console

import (
	"bufio"
	"github.com/wudaoluo/fyleaf/network"
)


type Agent struct {
	conn   *network.TCPConn
	reader *bufio.Reader
}
