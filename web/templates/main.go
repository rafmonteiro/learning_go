package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

var tpl *template.Template

// var fm = template.FuncMap{
// 	"uc": strings.ToUpper,
// 	"ft": firstThree,
// }

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
