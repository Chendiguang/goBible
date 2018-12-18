package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	// call
	forEachNode(doc, startElement, endElement)
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

//!+ startElement && endElement
var depth int
var stack []string // similar to stack
var needRight bool // try to ensure it should add ">" at the monment

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		l := len(stack)
		if l != 0 && stack[l-1] != n.Data && needRight {
			fmt.Println(">")
		}
		stack = append(stack, n.Data)
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		needRight = true
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		l := len(stack)
		if l != 0 && stack[l-1] == n.Data {
			fmt.Println("/>")
			if l > 1 {
				stack = stack[:l-2]
			}
		} else {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
			if l > 0 {
				stack = stack[:l-1]
			}
		}
		needRight = false
	}
}

//!- startElement && endElement
