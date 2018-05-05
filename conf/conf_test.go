package conf

import (
	"testing"
	"fmt"
	"time"
)



func Test_JsonConf(t *testing.T) {
	a := GetInstance()
	a.ParseConf("../cmd/server.json","json")
	fmt.Println(a.Cfg.Version)
	a.Cfg.Version = "1.121"
	a.Cfg.Mysql.DBpasswd = "aa"
	fmt.Println(a.Cfg.Version)

	a.C.SaveFile()
	fmt.Println(a.Cfg.Version)
	time.Sleep(10*time.Second)
	a.C.Reload()
	fmt.Println(a.Cfg.Version)
	t.Log(a.Cfg)
}



func Test_IniConf(t *testing.T) {
	a := GetInstance()
	a.ParseConf("../cmd/server.ini","ini")
	//fmt.Println(a.Cfg.Version)


	a.C.Update("","Version","1.141")

	a.C.SaveFile()
	//fmt.Println(a.Cfg.Version)
	//time.Sleep(10*time.Second)
	//a.C.Reload()
	//fmt.Println(a.Cfg.Version)
	//t.Log(a.Cfg)

}

