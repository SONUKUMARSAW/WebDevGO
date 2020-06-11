package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("temp/*.gohtml"))
}

type person struct{
	Name string
	Age int
	Contact string
}

func main(){
	//data := []string{"C++","Golang","Python","C","HTML"}
	//err:= tpl.ExecuteTemplate(os.Stdout,"01slice.gohtml",data)
	
	// data := map[string]string{
	// 	"name":"Sonu Kumar Saw",
	// 	"Age":"21",
	// 	"Address":"829150-Ramgarh, Jharkhand, India.",
	// }
	//err:= tpl.ExecuteTemplate(os.Stdout,"02map.gohtml",data)
	data := []person{
		person{"Sonu Kumar Saw",21,"sonukumarsaw66@gmail.com"},
		person{"James Bond",48,"jamebond@agent007.com"},

	}
	err:= tpl.ExecuteTemplate(os.Stdout,"03struct.gohtml",data)
	if err!=nil {
		log.Fatalln("error:",err)
	}
}