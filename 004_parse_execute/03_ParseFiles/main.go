package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	/*
	 Get a template to a file
	*/
	tpl, err := template.ParseFiles("one.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	/*
			 creates a new Template and parses the template definitions from the named files.
			 The returned template's name will have the base name and parsed contents of the first file

		 	tpl below has more than one template
	*/
	tpl, err = tpl.ParseFiles("two.gmao", "vespa.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	/*
		Since tpl has more than one template, you need to use ExecuteTemplate
		you need to pass the specific name of the template
		in ExecuteTemplate() function
	*/
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
