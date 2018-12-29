package main

import (
	"flag"
	"fmt"
	"goBible/chapter7/exam7-6/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature, e.g., \"100C\"")

// var k = tempconv.CelsiusFlag("k", 20.0, "The Kelvin temperature, e.g., \"100K\"")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
