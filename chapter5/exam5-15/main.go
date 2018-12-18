package main

import (
	"fmt"
	"math"
)

func max(vals ...int) int {
	res := math.MinInt32
	for _, val := range vals {
		if val > res {
			res = val
		}
	}
	return res
}

func min(vals ...int) int {
	res := math.MaxInt32
	for _, val := range vals {
		if val < res {
			res = val
		}
	}
	return res
}

func main() {
	fmt.Println(max())
	fmt.Println(max(1, 2))
	fmt.Println(max(1, 2, 3))
}
