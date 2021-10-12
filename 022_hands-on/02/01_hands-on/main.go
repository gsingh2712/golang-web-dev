package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// TCP Server initialise
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err) // Check log.Fatal Signature , it prints log to Std output and exits
	}
	defer l.Close() // This is done to close connection as control exits main function

	for {
		conn, err := l.Accept() // Important step , accepts incoming connection
		if err != nil {
			log.Println(err)
		}
		// Wrting responce to connection
		io.WriteString(conn, "Connection established\n") // Connection implements Writer interface
		conn.Close()
	}
}
