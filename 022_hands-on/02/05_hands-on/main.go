package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {

		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		scanner := bufio.NewScanner(conn)

		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			if ln == "" {
				// when ln is empty, header is done
				fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
				break
			}
		}
		fmt.Println("Code got here.")
		io.WriteString(conn, "Connection Established\n")
		conn.Close()
	}

}
