package conf

import (
	"sync"
	"github.com/Unknwon/goconfig"
	"github.com/wudaoluo/fyleaf/glog"
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
		glog.Fatal(errstr.F_OnceLoad,err)
	}

	err = i.parse()
	if err != nil {
		glog.Fatal(errstr.F_INIParse,err)
	}

	return

}


//解析cfg --> i.server
func (i *iniConf) parse() error {
	//开始解析配置文件
	i.mu.Lock()
	defer i.mu.Unlock()
	//glog.Info(errstr.I_Lock)
	//default
	i.server.Version = i.cfg.MustValue(goconfig.DEFAULT_SECTION,"Version","1.0")
	i.server.LogLevel =i.cfg.MustValue(goconfig.DEFAULT_SECTION,"LogLevel","FATAL")
	i.server.LogPath =i.cfg.MustValue(goconfig.DEFAULT_SECTION,"LogPath","log")
	i.server.WSAddr = i.cfg.MustValue(goconfig.DEFAULT_SECTION,"WSAddr","127.0.0.1:3653")
	i.server.TCPAddr = i.cfg.MustValue(goconfig.DEFAULT_SECTION,"TCPAddr","127.0.0.1:3654")
	i.server.MaxConnNum = int32(i.cfg.MustInt(goconfig.DEFAULT_SECTION,"MaxConnNum",100))
	i.server.ConsolePort = i.cfg.MustInt(goconfig.DEFAULT_SECTION,"ConsolePort",0)

	//mysql
	i.server.Mysql.DBname = i.cfg.MustValue("Mysql","DBname")
	i.server.Mysql.DBaddr = i.cfg.MustValue("Mysql","DBaddr")
	i.server.Mysql.DBport = i.cfg.MustValue("Mysql","DBport")
	i.server.Mysql.DBuser = i.cfg.MustValue("Mysql","DBuser")
	i.server.Mysql.DBpasswd = i.cfg.MustValue("Mysql","DBpasswd")


	//glog.Info(errstr.I_Unlock)
	return nil
}


func (i *iniConf) Reload() {
	err := i.cfg.Reload()
	if err !=nil {
		glog.Error(errstr.E_Reload,err)
	}

	//需要赋值

	return

}


func (i *iniConf) SaveFile() {
	err := goconfig.SaveConfigFile(i.cfg,i.fileName)
	if err != nil {
		glog.Error(errstr.E_SaveFaild,err)
	}
}


func (i *iniConf) Update(section,key,value string) {
	glog.Info(errstr.I_ModifyKey,section,key,value)
	i.mu.Lock()
	i.cfg.SetValue(section,key,value)
	i.mu.Unlock()
}


func (i *iniConf) Delete(section,key string) {
	ok := i.cfg.DeleteKey(section, key)
	if ok {
		glog.Info(errstr.I_DelKeySuccess)
	}else {
		glog.Error(errstr.I_DelKeyFaild)
	}
}
