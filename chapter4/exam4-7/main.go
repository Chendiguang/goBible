/*修改函数reverse,来翻转一个UTF-8编码的字符串中的字符元素,
 *传入参数是该字符串对应的字节slice类型([]byte).
 *你可以做到不需要重新分配内存就实现该功能吗
 * TO-DO
 */

package main

import (
	"fmt"
)

func main() {
	s := "Hello, 世界"
	t := []rune(s)
	fmt.Println(t)
	fmt.Println([]rune(s))
	fmt.Println(string([]byte{228}))
	reverse(&t)
	fmt.Println(string(t))
}

func reverse(s *[]rune) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		// fmt.Println(i, j)
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

// func reverseInplace(b []byte) []byte {

// }
