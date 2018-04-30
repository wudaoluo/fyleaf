package utils

import "fmt"

const  version = "fyleaf version:0.1"

func PrintVersion() {
	fmt.Println("version:", version)
}

func ReturnVersion() string{
	return version
}
