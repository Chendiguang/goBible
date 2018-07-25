package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// Listen 函数创建一个net.Listener对象
	port := flag.String("port", "8000", "http listen port")
	flag.Parse()
	addr := "localhost:" + *port
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		// handle one connection at a time.
		go handleConn(conn, addr)
	}
}

func handleConn(conn net.Conn, addr string) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, addr+" "+time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
