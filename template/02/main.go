package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main(){
	name := "Sonu Kumar Saw"
	htmlcode := `
	<html>
	<head><title>templates</title></head>
	<body>
	`+name+`
	</body>
	</html>
	`
	fd, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	
	_,err = io.Copy(fd,strings.NewReader(htmlcode))
	if err != nil {
		log.Fatalln(err)
	}
	defer  fd.Close()
}