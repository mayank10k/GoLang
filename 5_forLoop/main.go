package main

import "fmt"

//go has only for loop --> we wil implement while loop using for loop
func main(){

	//while loop
	// i:=1
	// for i<=3 {
	// 	fmt.Println(i);
	// 	i++
	// }
   

	// for{
	// 	// infinite loop
	// }


	//classic for looop

	for i:=0;i<=3;i++{
		if i==2 {
			continue
		}
		fmt.Println(i)
	}

	for j:= range 5{
		fmt.Println(j)
	}

}
