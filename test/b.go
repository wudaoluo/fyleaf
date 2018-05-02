package main

import (
	"reflect"
	"fmt"
)


type msg1111 struct{
	User string
	Pwd  string
}

func main() {


	msgType := reflect.TypeOf(&msg1111{})
	//fmt.Println(msgType.Elem().Name())
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		fmt.Println("json message pointer required")
		//return nil, errors.New("json message pointer required")
	}
	msgID := msgType.Elem().Name()

	fmt.Println(msgID)
}
