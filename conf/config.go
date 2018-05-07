package conf

import (
	"sync"
	"github.com/wudaoluo/fyleaf/glog"
)

type Config interface {
	load()  	           //加载配置文件
	Reload()               //从本地重新载入配置文件
	SaveFile()             //保存配置文件到本地
	Update(string,string,string)
}


//这个还没有完成，等写完console再说吧
type CfgModify interface {
	Update(string,string,string)   // section,key,value
	Delete(string,string)          // section,key   int:0, bool:false, string:""
}


//验证结构体是否实现接口
var (
	_ Config = new(iniConf)
	_ Config = new(jsonConf)

	_ CfgModify = new(iniConf)
	_ CfgModify = new(jsonConf)
)




type singleton struct {
	Cfg  	Server
	C       Config
	CM      CfgModify
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
		glog.Fatal(errstr.F_NotType,ftype)
	}

	s.C.load()
}