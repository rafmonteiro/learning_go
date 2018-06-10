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
	// 01 = month
	// 02 = day
	// 03 = hour
	// 04 = minutes
	// 05 = seconds
	// 06 = year
	// 07 = timezone
	// MST = I don't remember.
	return t.Format("01-02-2006 03:04 05 -0700 MST")
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
