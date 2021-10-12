package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	/*
	   Parsing the file into template and error checking with 'Must' call
	*/
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	/*
		Check tpl.gohtml, initialising wisdom by assinging it to '.'
	*/
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Release self-focus; embrace other-focus.`)
	if err != nil {
		log.Fatalln(err)
	}
}
