package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
func init()  {
	tpl = template.Must(template.ParseFiles("temp/index.gohtml"))
}

func main(){
	err := tpl.ExecuteTemplate(os.Stdout,"index.gohtml","Sonu Kumar Saw")
	if err != nil {
		log.Fatalln(err)
	}
}