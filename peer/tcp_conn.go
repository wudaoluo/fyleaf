package peer

import (
	"net"
	"sync"
)


//TODO  这个有啥鸟用
type ConnSet map[net.Conn]struct{}

type TCPConn struct {
	sync.Mutex
	conn      net.Conn
	writeChan chan []byte
	closeFlag bool
	msgParser *MsgParser
}

