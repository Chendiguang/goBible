/* Lissajous generates GIF animations of random Lissajous figures.
Uasge:
	go build goBible/chapter1/exam1-9
	./exam1-9 http://gopl.io http://nnn.bad baidu.com
	output:
		200 OK
		fetch: Get http://nnn.bad: dial tcp: lookup nnn.bad: no such host
		200 OK
*/
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			// os.Exit(1)
			continue
		}

		fmt.Printf("StatusCode: %s\n", resp.Status)
	}
}
