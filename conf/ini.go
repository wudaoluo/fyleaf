package conf

import (
	"sync"
	"github.com/Unknwon/goconfig"
	"fyleaf/glog"
)

type iniConf struct {
	mu      sync.Mutex
	server  *Server
	fileName string
	cfg      *goconfig.ConfigFile
}


func newiniConf(f string,cfg *Server) *iniConf {
	return &iniConf{
		fileName:f,
		server:cfg,
	}
}



func (i *iniConf) load() {
	var err error
	i.cfg, err = goconfig.LoadConfigFile(i.fileName)
	if err != nil {
		glog.Fatal("第一次载入配置文件失败",err)
	}

	err = i.parse()
	if err != nil {
		glog.Fatal("解析ini失败",err)
	}

	return

}


//解析cfg --> i.server
func (i *iniConf) parse() error {
	//开始解析配置文件
	i.mu.Lock()
	glog.Info("开始-修改配置文件加锁")
	//default
	i.server.Version = i.cfg.MustValue(goconfig.DEFAULT_SECTION,"Version","1.0")
	i.server.LogLevel =i.cfg.MustValue(goconfig.DEFAULT_SECTION,"LogLevel","FATAL")
	i.server.WSAddr = i.cfg.MustValue(goconfig.DEFAULT_SECTION,"WSAddr","127.0.0.1:3653")
	i.server.TCPAddr = i.cfg.MustValue(goconfig.DEFAULT_SECTION,"TCPAddr","127.0.0.1:3654")
	i.server.MaxConnNum = i.cfg.MustInt(goconfig.DEFAULT_SECTION,"MaxConnNum",20000)
	i.server.ConsolePort = i.cfg.MustInt(goconfig.DEFAULT_SECTION,"ConsolePort",7771)
	//
	////mysql
	i.server.Mysql.DBname = i.cfg.MustValue("Mysql","DBname")
	i.server.Mysql.DBaddr = i.cfg.MustValue("Mysql","DBaddr")
	i.server.Mysql.DBport = i.cfg.MustValue("Mysql","DBport")
	i.server.Mysql.DBuser = i.cfg.MustValue("Mysql","DBuser")
	i.server.Mysql.DBpasswd = i.cfg.MustValue("Mysql","DBpasswd")

	i.mu.Unlock()
	glog.Info("完成-修改配置文件解锁")
	return nil
}


func (i *iniConf) Reload() {
	err := i.cfg.Reload()
	if err !=nil {
		glog.Error("重新载入配置文件失败",err)
	}

	//需要赋值

	return

}


func (i *iniConf) SaveFile() {
	err := goconfig.SaveConfigFile(i.cfg,i.fileName)
	if err != nil {
		glog.Error("保存配置到本地失败",err)
	}
}


func (i *iniConf) Update(section,key,value string) {
	glog.Info("修改配置值",section,key,value)
	i.mu.Lock()
	i.cfg.SetValue(section,key,value)
	i.mu.Unlock()
}


func (i *iniConf) Delete(section,key string) {
	ok := i.cfg.DeleteKey(section, key)
	if ok {
		glog.Info("删除key成功")
	}else {
		glog.Error("key不存在 or 删除key失败")
	}
}
