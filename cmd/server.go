package main

import (
	"fyleaf/glog"
)

func main() {
	//flag.Parse()
	glog.Init("a","log",2)
	    // 1


	glog.Info("This is a Info log")         // 2
	glog.Warning("This is a Warning log")
	glog.Error("This is a Error log")

	glog.V(1).Infoln("level 1")     // 3
	glog.V(2).Infoln("level 2")

	glog.Flush()    // 4

}
