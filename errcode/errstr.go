package errcode

import "sync"

type utils struct {
	StopTimer string
}



type singleton struct {
	Utils  utils
}


var instance *singleton
var once sync.Once


func GetErrInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func (a *singleton) Load() {
	a.Utils.StopTimer = "停止时间轮训 timer"
}