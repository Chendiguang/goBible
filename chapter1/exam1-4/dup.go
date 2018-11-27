package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	res := []string{}
	// default reader is os.Stdin
	if len(files) == 0 {
		countLines(os.Stdin)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			if countLines(f) {
				fmt.Println(file)
			}
		}
	}

	for _, name := range res {
		fmt.Println(name)
	}
}

func countLines(f *os.File) bool {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		// should exit loop explicity, when input is os.Stdin
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			return true
		}
	}
	return false
}
