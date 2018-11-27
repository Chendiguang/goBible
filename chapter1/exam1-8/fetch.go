/* Lissajous generates GIF animations of random Lissajous figures.
Uasge:
	go build goBible/chapter1/exam1-8
	./exam1-8 http://gopl.io http://nnn.bad baidu.com
*/
package main

import (
	"fmt"
	"io"
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
		defer resp.Body.Close()
		_, err = io.Copy(os.Stdout, resp.Body)
		// data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// fmt.Printf("%s\n", data)
	}
}
