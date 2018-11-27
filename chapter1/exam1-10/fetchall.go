/*
./fetchall https://www.lagou.com
1st:
	0.83s  463972  https://www.lagou.com
	0.83s elapsed
2nd:
	0.77s  440328  https://www.lagou.com
	0.77s elapsed
*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Fetchall fetch URLs in parallel and reports their times and sizes.
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// if !strings.HasPrefix(url, "http://") {
		// 	url = "http://" + url
		// }
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	delta := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d  %s", delta, nbytes, url)
}
