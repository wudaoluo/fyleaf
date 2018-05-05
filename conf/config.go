package conf

import (
	"sync"
	"fyleaf/glog"
)

type Config interface {
	load()  	           //加载配置文件
	Reload()               //从本地重新载入配置文件
	SaveFile()             //保存配置文件到本地
}





type Server struct {
	Version     string
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
	Debug       bool
	Mysql       *MysqlInfo
}

//root:123456@tcp(10.211.55.4:3306)/game
type MysqlInfo struct {
	DBname string
	DBaddr string
	DBport string
	DBuser string
	DBpasswd string
}



type singleton struct {
	Cfg  	Server
	C        Config
}


var instance *singleton
var once sync.Once


func  GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}



func (s *singleton)ParseConf(f,ftype string) {
	switch ftype{
	case "json":
		s.C = newjsonConf(f,&s.Cfg)
	case "ini":
		s.C =newiniConf(f,&s.Cfg)
	default:
		glog.Fatal("不支持类型格式",ftype)
	}

	s.C.load()
}