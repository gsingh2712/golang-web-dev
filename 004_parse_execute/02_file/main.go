package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml") // Create a template out of a file
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}
	defer nf.Close()

	/*
		https://pkg.go.dev/text/template#Template.Execute
		Execute applies a parsed template to the specified data object, and writes the output to wr
		func (t *Template) Execute(wr io.Writer, data interface{}) error
	*/
	err = tpl.Execute(nf, nil) // It writes the result to index.html in this case
	if err != nil {
		log.Fatalln(err)
	}
}

// Do not use the above code in production
// We will learn about efficiency improvements soon!
