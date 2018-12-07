/*
 * iota 机制的存在局限性，例如:
 * 因为不存在指数运算符，所以无从生成更为人熟知的1000的幂。
 *
 * 要求：用尽可能简洁的方法声明从KB、MB、直到YB的常量。
 * 1KB = 1000Byte
 * 1KiB = 1024Byte
 */
package main

import (
	"fmt"
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)
const (
	KB, MB, GB, TB, PB, EB, ZB, YB = 1000, KB * KB, MB * KB, GB * KB, TB * GB, PB * KB, EB * KB, ZB * KB
)

func main() {
	fmt.Println(KB, MB, GB, TB)
}
