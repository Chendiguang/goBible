package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
)

func main() {
	v := flag.String("value", "x", "set the translation type")
	c := sha256.Sum256([]byte(*v))
	res := PopCount(8, c)
	fmt.Println(res)
}

func PopCount(x uint64, pc [32]byte) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
