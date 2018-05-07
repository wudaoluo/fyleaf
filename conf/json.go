package conf

import (
	"github.com/wudaoluo/fyleaf/glog"
	"io/ioutil"
	"sync"
	"github.com/json-iterator/go"
	"os"
)

type jsonConf struct {
	mu      sync.Mutex
	server  *Server
	fileName string
}


func newjsonConf(f string,cfg *Server) *jsonConf {
	return &jsonConf{
		fileName:f,
		server:cfg,
	}
}


func (s *jsonConf) load() {
	s.mu.Lock()
	defer s.mu.Unlock()
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		panic(errstr.F_OnceLoad)

	}

	err = jsoniter.Unmarshal(data, &s.server)

	if err != nil {
		panic(errstr.F_JsonParse)

	}

	s.defaultValue()
}



func (s *jsonConf) Reload() {
	glog.Info(errstr.I_ReloadConf)
	s.mu.Lock()
	defer s.mu.Unlock()
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		glog.Error(errstr.E_OnceLoad)
	}

	err = jsoniter.Unmarshal(data, &s.server)

	if err != nil {
		glog.Error(errstr.E_JsonParse)

	}

	s.defaultValue()

}

/*
	i.server.Version = i.cfg.MustValue(goconfig.DEFAULT_SECTION,"Version","1.0")
	i.server.LogLevel =i.cfg.MustValue(goconfig.DEFAULT_SECTION,"LogLevel","FATAL")
	i.server.LogPath =i.cfg.MustValue(goconfig.DEFAULT_SECTION,"LogPath","log")
	i.server.WSAddr = i.cfg.MustValue(goconfig.DEFAULT_SECTION,"WSAddr","127.0.0.1:3653")
	i.server.TCPAddr = i.cfg.MustValue(goconfig.DEFAULT_SECTION,"TCPAddr","127.0.0.1:3654")
	i.server.MaxConnNum = i.cfg.MustInt(goconfig.DEFAULT_SECTION,"MaxConnNum",100)
	i.server.ConsolePort = i.cfg.MustInt(goconfig.DEFAULT_SECTION,"ConsolePort",0)

*/
func (s *jsonConf) defaultValue() {
	if s.server.Version == "" {
		s.server.Version = "1.0"
	}

	if s.server.LogLevel == "" {
		s.server.LogLevel = "FATAL"
	}

	if s.server.LogPath == "" {
		s.server.LogPath = "log"
	}

	if s.server.WSAddr == "" {
		s.server.WSAddr = "127.0.0.1:3653"
	}

	if s.server.TCPAddr == "" {
		s.server.TCPAddr = "127.0.0.1:3654"
	}

	if s.server.MaxConnNum <= 0 {
		s.server.MaxConnNum = 100
	}
}


/*
1. 解析json，失败返回
2. 创建临时文件，失败返回
3. 写入临时文件，失败返回
4. rename文件， 失败返回
5. 删除临时文件
*/
func (s *jsonConf) SaveFile() {
	var err error

	s.mu.Lock()
	defer s.mu.Unlock()

	saveData,err := jsoniter.MarshalIndent(s.server,"","    ")
	if err != nil {
		glog.Error(errstr.E_JsonParse)
		return
	}

	temp, err := ioutil.TempFile("/tmp", ".tmp")
	if err != nil {
		glog.Error(errstr.E_CreateFaild,err)
		return
	}
	defer temp.Close()

	_, err = temp.Write(saveData)
	if err != nil {
		glog.Error(errstr.E_WriteFileFaild,err)
		return
	}

	defer os.Remove(temp.Name())

	//改变文件名称
	err = os.Rename(temp.Name(),s.fileName)

	if err != nil {
		glog.Error(errstr.E_CopyConfFaild)
		return
	}

	glog.Info(errstr.I_CopyConfSuccess)
}




func (i *jsonConf) Update(section,key,value string) {

}


func (i *jsonConf) Delete(section,key string) {

}