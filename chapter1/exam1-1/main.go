package main

import (
	"fmt"
	"os"
)

// Echo prints its command-line name.
func main() {
	fmt.Println(os.Args[0])
}
