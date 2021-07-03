package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	/*
	  Writing different words/output to ResponseWriter depending on Path
	*/
	case "/dog":
		io.WriteString(w, "doggy doggy doggy")
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")
	}

	/*
		Not an eleganat way for Path based routing handling
		Check 02_NewServeMux for the same.
	*/

}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
