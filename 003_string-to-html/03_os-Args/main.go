package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
		</head>
		<body>
		<h1>` +
		name +
		`</h1>
		</body>
		</html>
	`)

	nf, err := os.Create("index.html") // Returns a Pointer to File created
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	/*
		strings.NewReader(str) returns a new Reader reading from 'str'
	*/
	io.Copy(nf, strings.NewReader(str))
}

/*
at the terminal:
1 go run main.go Gary
2 open index.html

*/
