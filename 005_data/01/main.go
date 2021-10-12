package main

import (
	"log"
	"os"
	"text/template"
)

/*
Shows how to pass data to a template to File containing  Placeholder --> {{.}}

Dot here refers to the current value of the Data at that point of time in execution

*/

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
