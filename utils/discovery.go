package utils


import (
	consulapi "github.com/hashicorp/consul/api"
	"fmt"
)


type Discovery interface {
	Run()
	Stop()
}


type discovery struct {

	D Discovery
}


func Newdiscovery() {

}




type consulStruct struct {
	name string
	addr string
	port int
	tag  []string
	consul_addr string

}



func NewConsul(name,addr,consul_addr string,port int,tag []string) *consulStruct{
	return &consulStruct{
		name:name,
		addr:addr,
		port:port,
		tag:tag,
		consul_addr:consul_addr,
	}
}

func (cs *consulStruct) ServDiscover() error{
	var err error
	config := consulapi.DefaultConfig()
	config.Address = cs.consul_addr
	client, err := consulapi.NewClient(config)
	if err != nil {
		return err
	}

	//创建一个新服务。
	reg := new(consulapi.AgentServiceRegistration)
	reg.Name = cs.name
	reg.Address = cs.addr
	reg.Port = cs.port
	reg.Tags = cs.tag

	//增加check
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d%s", reg.Address, reg.Port, "/check")
	//设置超时 5s。
	check.Timeout = "5s"
	//设置间隔 5s。
	check.Interval = "5s"

	reg.Check = check

	err = client.Agent().ServiceRegister(reg)
	return err
}



func httprequset() {
	fmt.Println("版本发现")
}
