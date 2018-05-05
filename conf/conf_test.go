package conf

import (
	"testing"
	"fmt"
)



func Test_JsonConf(t *testing.T) {
	a := GetInstance()
	a.ParseConf("../cmd/server.json","json")
	fmt.Println(a.Cfg.Version)
	a.Cfg.Version = "1.0"
	fmt.Println(a.Cfg.Version)
	t.Log(a.Cfg)
}
