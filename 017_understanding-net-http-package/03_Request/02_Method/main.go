package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() // Parse the request
	if err != nil {
		log.Fatalln(err)
	}

	/*
		Creating and Struct and assigning it values post parsing the request
	*/
	data := struct {
		Method      string     // Check how these are referenced in index.html
		Submissions url.Values // Any reason? First Letter member variables in caps!! -> so that they can be exported and accessed
	}{
		req.Method,
		req.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
