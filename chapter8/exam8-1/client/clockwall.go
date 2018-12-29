/*
 * ./clockwall NewYork=localhost:8010 London=localhost:8011 Tokyo=localhost:8012
 */

package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// flag.Parse()
	for _, item := range os.Args[1:] {
		addrs := strings.Split(item, "=")
		go handleClient(addrs[1])
	}
	select {}
}

func handleClient(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Fatal(err)
	}
}
