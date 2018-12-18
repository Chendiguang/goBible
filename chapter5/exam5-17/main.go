/*
 * Usage:
 * 		go run goBible/chapter5/exam5-17/main.go http://gopl.io
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

func main() {
	defer trace("main")() // optional
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		doc, err := html.Parse(resp.Body)
		if err != nil {
			log.Println("html parse error.")
			continue
		}
		res := ElementsByTagName(doc, "a", "meta")
		fmt.Println(res)
		// for _, n := range res {
		// 	fmt.Printf("node: %v, node.Data: %s\n", n, n.Data)
		// }
	}
}

// time tests
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	if len(name) == 0 {
		return nil
	}
	var visit func(nodes []*html.Node, n *html.Node) []*html.Node
	visit = func(nodes []*html.Node, n *html.Node) []*html.Node {
		if n.Type == html.ElementNode {
			for _, tag := range name {
				if n.Data == tag {
					nodes = append(nodes, n)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			nodes = visit(nodes, c)
		}
		return nodes
	}
	return visit(nil, doc)
}
