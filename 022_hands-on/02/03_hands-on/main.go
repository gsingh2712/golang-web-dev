package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Open a tcp server on 8080
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err) // Print Error and exits
	}

	defer l.Close() // Close the TCP

	for { // infinite Loop till it accepts connection
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		scanner := bufio.NewScanner(conn)

		for scanner.Scan() { // check if there is a incoming line present
			line := scanner.Text() // read the current line
			fmt.Println(line)
		}

		defer conn.Close() // defered Closing Connection so that it is closed
		// after for loop exits since we are wriing to connection as well

		fmt.Println("Code Flow reached here") // This line is only seen when you close localhost:8080 in browser
		// since we got into loop with scanner.Scan() as it is scanning an inifinite stream

		io.WriteString(conn, "Connection established\n")
		// This line is not seen why? , because closing browser closes connection

	}
}
