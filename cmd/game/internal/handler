package internal

import (
	"reflect"
	"server/msg"
	"github.com/name5566/leaf/log"
	"server/model"
	"fmt"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.UserLogin{}, handleLogin)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
//
//func handleLogin(args []interface{}) {
//	// 收到的 Hello 消息
//	m := args[0].(*msg.UserLogin)
//	// 消息的发送者
//	a := args[1].(gate.Agent)
//
//	// 输出收到的消息的内容
//	log.Debug("hello %v", m.LoginName)
//	log.Debug(m.LoginPW)
//
//	// 给发送者回应一个 Hello 消息
//	a.WriteMsg(&msg.UserLogin{
//		LoginName: "client",
//	})
//}


func handleLogin(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.UserLogin)
	// 消息的发送者
	//a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("username %v", m.LoginName)
	log.Debug("password %v", m.LoginPW)


	aa:=&model.User{User:m.LoginName,Passwd:m.LoginPW}
	err:=aa.UserGetByName()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(aa.Id)
	if aa.Id == 0 {
		log.Debug("添加新的用户")
		id,err:=aa.UserAdd()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id)
	}else {
		log.Debug("用户已经存在")
	}

}