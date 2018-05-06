package utils

import (
	"fmt"
)

const  version = "0.1"

func PrintVersion() {
	fmt.Println("fyleaf version:", version)
}

func ReturnVersion() string{
	return version
}
