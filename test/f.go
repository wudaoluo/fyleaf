package main

import (
	"strings"
	"fmt"
)

func main() {
	var s string = "aaaaaaaa\r\n"
	fmt.Println(s[:len(s)-1])
	s=strings.TrimSuffix(s[:len(s)-1],"\r")
	fmt.Println(s)
}
