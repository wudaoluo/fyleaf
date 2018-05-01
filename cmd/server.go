package main

import (
	"fyleaf/glog"
	"fmt"
	"fyleaf/models/mongodb"
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


	//for {
	//	time.Sleep(10*time.Second)
	//	newfunc.GetInfo()
	//	os.Exit(0)
	//}

	c, err := mongodb.Dial("10.211.55.4", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()




}
