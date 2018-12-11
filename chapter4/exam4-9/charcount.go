//!+

// wordfreq 汇总输入文件中每个单词出现的次数。
// 在第一次调用Scan之前, 需要使用input.Split(bufio.ScanWords)来
// 将文本行按照单词分割, 二不是行分割。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// 初始化分割的格式, 可以按字/行/Unicode字符/单词来分割
	// bufio.ScanBytes
	// bufio.ScanLines
	// bufio.ScanRunes
	// bufio.ScanWords
	input.Split(bufio.ScanWords) // 必须在input.Scan()之前设定好
	for input.Scan() {
		words[input.Text()]++
	}
	defer func() {
		fmt.Printf("words\tcount\n")
		for c, n := range words {
			fmt.Printf("%q\t%d\n", c, n)
		}
	}()
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %q\n", err)
		os.Exit(1)
	}
}

//!-
