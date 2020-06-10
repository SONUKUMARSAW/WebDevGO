package main

import (
	"fmt"
)

func main() {
	name := "Sonu Kumar Saw"
	code := `
	<html>
	<head><title> Thats My Name </title>
	</head>
	<body>
	<h1>`+ name +`</h1>
	</body>
	</html>
	`
	fmt.Println(code)
}

// use go run main.go to print the html in console
// use go run main.go > index.html to print in file

