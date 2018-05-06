### glog
日志级别 INFO,WARNING,ERROR,FATAL


### 改后版本使用方式
可以自己创建目录
通过glog.Init()初始化
修改日志格式  I0501 12:57:47 server.go:13;This is a Info log
让所有日志都写入到一个文件


注意：为了让标准输出尽量少的输出 日志级别只能设置 FATAL
    //日志级别，日志目录，v
    glog.Init("a","log",2)

    glog.Info("This is a Info log")
    glog.Warning("This is a Warning log")
    glog.Error("This is a Error log")

    glog.V(1).Infoln("level 1")
    glog.V(2).Infoln("level 2")

    glog.Flush()


### 官方版本使用方式
    flag.Parse()    // 1

    glog.Info("This is a Info log")         // 2
    glog.Warning("This is a Warning log")
    glog.Error("This is a Error log")

    glog.V(1).Infoln("level 1")     // 3
    glog.V(2).Infoln("level 2")
    glog.Flush()    // 4
    go run server.go -log_dir="log" -v=2  -stderrthreshold="FATAL"
    比如 -v=3  就会输出 v(1),v(2),v(3)的日志,是一个很不错的功能

