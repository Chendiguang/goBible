/*
 * 一个变长版本的strings.Join()函数
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(join("_", "a", "b", "c"))
}

func join(seq string, vals ...string) string {
	// res := ""
	// for _, str := range vals {
	// 	strings.Join()
	// }
	return strings.Join(vals, seq)
}
