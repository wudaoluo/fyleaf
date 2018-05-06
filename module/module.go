package module

import (
	"sync"
	"runtime"
	"fyleaf/glog"
	"fyleaf/conf"
)

type Module interface {
	OnInit()
	OnDestroy()
	Run(closeSig chan bool)
}



type module struct {
	mi 			Module
	closeSig 	chan bool
	wg 			sync.WaitGroup //在Run() 没有退出的时候，一直阻塞 不会运行destroy()
}


var mods []*module


func Register(mi Module) {
	m := new(module)
	m.mi = mi
	m.closeSig = make(chan bool,1)

	mods = append(mods,m)
}


func Init() {
	for i:=0; i< len(mods); i++ {
		mods[i].mi.OnInit()
	}

	for i:=0; i < len(mods); i++ {
		m := mods[i]
		m.wg.Add(1)
		go run(m)
	}
}


func run(m *module) {
	m.mi.Run(m.closeSig)
	m.wg.Done()
}

//倒叙方式关闭module
func Destroy() {
	for i := len(mods) - 1; i >= 0; i-- {
		m := mods[i]
		m.closeSig <- true
		m.wg.Wait()
		destroy(m)
	}
}


func destroy(m *module) {
	defer func() {
		if r := recover(); r != nil {
			if conf.LenStackBuf > 0 {
				buf := make([]byte, conf.LenStackBuf)
				l := runtime.Stack(buf, false)
				glog.Error("%v: %s", r, buf[:l])
			} else {
				glog.Error("%v", r)
			}
		}
	}()
	m.mi.OnDestroy()
}