package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	/*
		How to handle path based routing using Mux

	*/

	mux := http.NewServeMux()
	// Different handlers for different path
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	/*
		mux is passed to Listen and Serve since it implements handler interface
		meaning it implements ServeHTTP function
	*/
	http.ListenAndServe(":8080", mux)
}
