package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen(":8080")
	if err != nil {
		log.Println("something wrong!!!!")
	}
	for {
		conn, _ := listener.Accept()
		go func(conn net.Conn) {

		}(conn)
	}
}
