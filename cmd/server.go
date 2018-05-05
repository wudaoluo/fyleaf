package main

import (
	"fyleaf/glog"

	"time"
	"fyleaf/utils"
	"fyleaf/conf"
	"fmt"
)

func main() {
	//flag.Parse()
	glog.Init("INFO","log",2)
	    // 1


	glog.Info("This is a Info log")         // 2
	glog.Warning("This is a Warning log")
	glog.Error("This is a Error log")

	glog.V(1).Infoln("level 1")     // 3
	glog.V(2).Infoln("level 2")

	glog.Flush()    // 4

	utils.NewGlogClear("log","log")


	a := conf.GetInstance()
	a.ParseConf("server.json","json")
	//a.C.Load()
	go a.C.Reload()
	//time.Sleep(time.Second*10)
	fmt.Println(a.Cfg.Version)
	a.Cfg.Version = "1.0"
	//fmt.Println(&a.Cfg.Version)




	//for {
	//	time.Sleep(10*time.Second)
	//	newfunc.GetInfo()
	//	os.Exit(0)
	//}
	//
	//c, err := mongodb.Dial("10.211.55.4", 100)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer c.Close()


	time.Sleep(1*time.Hour)



}
