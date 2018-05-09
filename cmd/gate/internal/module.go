package internal

import (
	"github.com/wudaoluo/fyleaf/gate"
	"github.com/wudaoluo/fyleaf/network"
)

//TODO 这里很奇怪,为什么 *gate.Gate 就报错
/*
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1284149]

*/

type Module struct {
	gate.Gate
}


func (m *Module) OnInit() {
	m.TCPServer = &network.TCPServer{}
}