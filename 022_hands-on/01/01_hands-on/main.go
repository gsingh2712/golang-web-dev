package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog")
}

func me(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "me")
}

func home(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "home")
}

func main() {

	/*
	 HandleFunc expects a func with inputs (ResponseWriter , Request)
	*/
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil) // Use nil to use internal DefaultMux
}
