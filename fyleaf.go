package fyleaf

import (
	"fyleaf/module"
	"os"
	"os/signal"
	"syscall"
	"fyleaf/glog"
	"fyleaf/utils"
	"fyleaf/console"
	"fyleaf/conf"
	"flag"
	"fmt"
)



type argStruct struct {
	version bool
	configfile string
	configtype string
	logv    int
}

var arg = new(argStruct)

func init() {
	flag.BoolVar(&arg.version,"version",false,"print version")
	flag.StringVar(&arg.configfile,"c","fyleaf.ini","specify config file")
	flag.StringVar(&arg.configtype,"type","ini","ini|json")
	flag.IntVar(&arg.logv,"V",3,"glog.V的级别")
}


//负责整个框架初始化
func Run(mods ...module.Module) {
	if !flag.Parsed() {
		panic("flag还没有初始化，请在入口添加flag.Parse()")
	}

	if arg.version {
		utils.PrintVersion()
		os.Exit(0)
	}


	// 配置文件初始化
	cfg := conf.GetInstance()
	cfg.ParseConf(arg.configfile,arg.configtype)

	//日志初始化
	glog.Init(cfg.Cfg.LogLevel,cfg.Cfg.LogPath,glog.Level(arg.logv))
	defer glog.Flush()

	//start
	glog.Info(errstr.I_StartInfo)

	//日志清理初始化  NewGlogClear(path string,t ...time.Duration)
	//TODO 以后改成Init形式
	utils.NewGlogClear(cfg.Cfg.LogPath)

	// module
	for i := 0; i < len(mods); i++ {
		module.Register(mods[i])
	}
	module.Init()

	// console
	console.Init()

	waitSignal()

	// close
	console.Destroy()
	module.Destroy()


}

//阻塞，只有执行信号才执行
func waitSignal() {
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM)
	sig := <-osSignals
	glog.Info(errstr.I_StopInfo, sig)
}



var errstr struct{
	I_StartInfo string
	I_StopInfo  string
}


func init() {
	errstr.I_StartInfo = fmt.Sprintf("fyleaf %v starting up",utils.ReturnVersion())
	errstr.I_StopInfo = "Leaf closing down signal:"
}




