/*
 * Usage: go build goBible/chapter5/exam5-3
 * ./fetch http://www.baidu.com | ./exam5-3
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

	visit(nil, doc)
}

func visit(links []string, n *html.Node) {
	// visit appends to links each link found in n
	if n.Type == html.TextNode {
		if n.Data != "" {
			fmt.Println(n.Data)
		}
	}
	// Recursive
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if n.Data == "script" || n.Data == "style" {
			continue
		}
		visit(links, c)
	}
}
