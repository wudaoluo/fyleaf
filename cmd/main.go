package main

import (
	"flag"
	"github.com/wudaoluo/fyleaf"
	"github.com/wudaoluo/fyleaf/cmd/gate"
)






func main() {
	flag.Parse()
	fyleaf.Run(gate.Module)

}