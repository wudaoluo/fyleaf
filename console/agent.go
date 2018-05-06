package console

import (
	"bufio"
	"fyleaf/network"
)


type Agent struct {
	conn   *network.TCPConn
	reader *bufio.Reader
}
