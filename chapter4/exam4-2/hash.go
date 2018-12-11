package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var Type string
var value string

func init() {
	flag.StringVar(&Type, "t", "sha256", "set the translation type")
	flag.StringVar(&value, "v", "", "set the value")
}

func main() {
	flag.Parse()
	switch Type {
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(value)))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(value)))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(value)))
	}
}
