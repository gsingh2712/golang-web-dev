package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog")
}

func me(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "me")
}

func home(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("base.html")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "base.html", "Gaurav")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
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
