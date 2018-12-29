// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		res := request(url)
		fmt.Println(res)
	}
}

//!-

func mirroredQuery(urls []string) string {
	reponse := make(chan string, 3)
	for _, url := range urls {
		go func(url string) {
			reponse <- request(url)
		}(url)
	}
	select {
	case x := <-reponse:
		close(reponse)
		return x
	}
}

func request(url string) (string, error) {
	// resp, err := http.Get(url)
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(reqest)
	if err != nil {
		return nil, fmt.Sprintf("fetch: %v\n", err)
		// os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Sprintf("fetch: reading %s: %v\n", url, err)
		// os.Exit(1)
	}
	return fmt.Sprintf("%s", b)
}
