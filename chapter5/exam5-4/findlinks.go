package main

//TO_DO
import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// prints the links in an HTML document read from standard input.
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	// visit(nil, doc)
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	// visit appends to links each link found in n
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	// Recursive
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
	// return visit(visit(links, n.FirstChild), n.NextSibling)
}
