package main

import (
	"fmt"
	"time"
)

func main() {

	//simple swicth

	//we dont need to write "break" in swich-case in go
	// i := 3

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("others")
	// }


	//multiple condition switch

	switch time.Now().Weekday() {
	case time.Sunday,time.Saturday:
		fmt.Println("its weekend")
	default:
		fmt.Println("its workday")

	}


	//type switch
	whoAmI:= func (i interface{}){   // i interface{} --> i can be of any type
		switch i.(type){
		case int:
			fmt.Println("type is integer")
		case string:
			fmt.Println("type is string")
		case bool:
			fmt.Println("type is boolean")
		default:
			fmt.Println("other")
		}

	}

	whoAmI("mayank");
	whoAmI(1)
	whoAmI(33.8)


}
