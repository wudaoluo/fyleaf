package conf

import (
	"fyleaf/glog"
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
		glog.Fatal("第一次载入配置文件失败",err)
	}

	err = jsoniter.Unmarshal(data, &s.server)

	if err != nil {
		glog.Fatal("解析配置文件格式错误",err)
	}
}


func (s *jsonConf) Reload() {
	glog.Info("reload 配置文件")
	s.load()
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
		glog.Error("json解析失败")
		return
	}

	temp, err := ioutil.TempFile("/tmp", ".tmp")
	if err != nil {
		glog.Error("临时文件创建错误",err)
		return
	}
	defer temp.Close()

	_, err = temp.Write(saveData)
	if err != nil {
		glog.Error("配置文件写入临时文件错误",err)
		return
	}

	defer os.Remove(temp.Name())

	//改变文件名称
	err = os.Rename(temp.Name(),s.fileName)

	if err != nil {
		glog.Error("copy 配置文件错误",err)
		return
	}

	glog.Info("copy配置文件成功")
}




func (i *jsonConf) Update(section,key,value string) {

}


func (i *jsonConf) Delete(section,key string) {

}