package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	/*
		Check tpl.gohtml to find how it parses '.'
	*/

	sages := []string{"Krishna", "Buddha", "Jesus", "Savarkar", "MLK"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
