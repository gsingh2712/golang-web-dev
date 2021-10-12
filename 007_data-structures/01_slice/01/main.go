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

/*
	go run main.go > out.html
	open out.html
*/

/*
	Check tpl.gohtml , how it is parsing slice
	it is a bit confusing
	'.' in  {{range .}} ---> refers to slice
	while '.' in  {{.}} ---> refers to elements in slice
*/
func main() {

	sages := []string{"Krishna", "Buddha", "Jesus", "Savarkar", "MLK"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
