// Clock is a TCP server that periodically writes the time.
// TZ=US/Eastern ./clock -port 8010 &
// TZ=Asia/Tokyo ./clock -port 8012 &
// TZ=Europe/London ./clock -port 8011 &

package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	port := flag.String("port", "8000", "http listen port")
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}
