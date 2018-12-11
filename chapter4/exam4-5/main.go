package main

import (
	"fmt"
)

func main() {
	s := []string{"1", "1", "2", "2", "2", "3", "2", "4", "4"}
	fmt.Println(removeDulicate(s))
}

// Double pointer
func removeDulicate(s []string) []string {
	if len(s) < 2 {
		return s
	}
	i := 0
	for j := 1; j < len(s); {
		if s[i] == s[j] {
			j++
		} else {
			i++
			s[i] = s[j]
		}
	}
	return s[:i+1]
}
