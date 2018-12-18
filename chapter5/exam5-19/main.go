package main

import (
	"fmt"
)

func main() {
	fmt.Println(RcoverAndPanic())
}

func RcoverAndPanic() (ret int) {
	defer func() {
		if p := recover(); p != nil {
			ret = 1
		} else {
			ret = 2
		}
	}()
	panic(ret)
}
