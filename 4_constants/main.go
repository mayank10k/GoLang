package main

import "fmt"

const age = 10

// name:="maynak"  --> this is not allowed outside the function
func main() {
	//if declared once cant be redeclared again
	const name = "mayank golang"

	fmt.Println(age)

	
	const(
		port=5000
		host="localhost"
	)

	fmt.Println(port,host)
	



}