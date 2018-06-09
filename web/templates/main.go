package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	//tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

type sage struct {
	Name  string
	Motto string
}

func main() {

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", g)
	if err != nil {
		log.Fatalln(err)
	}
}
