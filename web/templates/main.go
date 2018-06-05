package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sages := map[string]string{
		"India":    "Gandhi",
		"USA":      "MLK",
		"Medidate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Morpheus",
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
