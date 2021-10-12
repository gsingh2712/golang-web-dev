package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	/*
		func Must(t *Template, err error) *Template
		Must is a helper that wraps a call to a function returning (*Template, error)
		and panics if the error is non-nil. It is intended for use in variable initializations
	*/
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	/*
		One the templates will be written to stdout
	*/
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	/*
		One the templates will be written to stdout
	*/
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
