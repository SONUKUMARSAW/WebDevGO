package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("temp/index.gohtml"))
}

func main()  {
	fd,err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(fd,"index.gohtml",os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer fd.Close()

}