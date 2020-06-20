package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		io.WriteString(conn, "Hello from the TCP\n")
		fmt.Fprintf(conn, "You are connected to the port :8080\n")
		fmt.Fprintf(conn, "Connection is about to close ...")
		conn.Close()
	}
}
