package main

import (
	"fmt"
	"os"
)

// Echo prints its command-line name.
func main() {
	for idx, name := range os.Args[1:] {
		fmt.Printf("Args of command-line: index %d, name %s\n",
			idx, name)
	}
}
