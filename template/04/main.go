package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("temp/*.gohtml"))
}
func main(){

	err :=tpl.ExecuteTemplate(os.Stdout,"first.gohtml","Sonu")
	check(err)
	err = tpl.ExecuteTemplate(os.Stdout,"middle.gohtml","Kumar")
	check(err)
	err = tpl.ExecuteTemplate(os.Stdout,"last.gohtml","Saw")
	check(err)
}

func check(err error){
	if err != nil {
		log.Fatalln("error :",err)
	}
}