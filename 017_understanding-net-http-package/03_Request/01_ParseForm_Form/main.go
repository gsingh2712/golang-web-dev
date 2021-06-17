package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

/*
Crux for http Server in Golang
   1. is you need handler interface => ServeHTTP( http.ResponseWriter, *http.Request)
   2. Use http.ListenAndServe(route , handler)
		To Handle any Request on the route
*/
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	/*
	 Read Request Struct in http package
	 Calling ParseForm() is compulsory to parse Data from Server and populate 'Postform' and 'Form'
	 Read about member Variables 'PostForm' and 'Form'
	*/
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form) // Form in Request struct is actually a map[string][]string (map to a slice of string)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
