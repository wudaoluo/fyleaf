package conf

import (
	"fyleaf/glog"
	"io/ioutil"
	"sync"
	"github.com/json-iterator/go"
	"encoding/json"
	"time"
	"fmt"
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
		glog.Fatal("第一次载入配置文件失败",err)
	}

	jsoniter.Unmarshal(data, &s.server)

	err = json.Unmarshal(data, &s.server)
	if err != nil {
		glog.Fatal("解析配置文件格式错误",err)
	}
	//fmt.Println(&s.server)
}


func (s *jsonConf) Reload() {
	for {
		time.Sleep(time.Second*1)
		fmt.Println(s.server.Version)
	}
}


func (s *jsonConf) SaveFile() {

}
