/*
 * Usage: go build goBible/chapter5/exam5-2
 * ./fetch http://www.baidu.com | ./exam5-2
 */
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "count: %v\n", err)
		os.Exit(1)
	}
	mp := make(map[string]int)
	count(mp, doc)
	for element, num := range mp {
		fmt.Printf("%s\t%d\n", element, num)
	}
}

func count(mp map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		mp[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count(mp, c)
	}
}
