package echo

import (
	"fmt"
	"strings"
)

// Use "+" and var
func echo1(args []string) {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

// Use "+" and Short statement
func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// Use strings.Join
func echo3(args []string) {
	fmt.Println(strings.Join(args, " "))
}
