package main

import "fmt"

func main() {
	// age := 17

	// if age >= 18 {
	// 	fmt.Println("person is adult")
	// }else if age>=12{
	// 	fmt.Println("person is teenager")
	// }else{
	// 	fmt.Println("person is a kid")
	// }

	
	var role="admin"
	var hasPermission=false

	if role=="admin" && hasPermission{
		fmt.Println("you are allowed")
	}else{
		fmt.Println("you are not allowed")
	}


    // we can declare a variable iside a if construct

	if age:=17; age>=18{
		fmt.Println("person is an adult")
	}else if age>12{
		fmt.Println("person is a teenager")
	}else{
		fmt.Println("person is a kid")
	}

	//go doesn't have ternary ,you will have to use normal if else
}
