package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	for _, item := range flag.Args() {
		addr := strings.Split(item, "=")[1]
		go helper(addr)
	}
	// helper("localhost:8010")
	for {

	}
}

func helper(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
