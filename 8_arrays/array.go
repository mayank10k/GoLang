package main

import "fmt"

func main() {

	//zeroed value:
	// int-->0  ,string --> " " ,bool-->false

	// var nums [4]int

	// nums[0]=1
	// fmt.Println(nums[0])

	// fmt.Println(nums)

	//array length
	// fmt.Println(len(nums))

	var vals [4]bool
	fmt.Println(vals)

	var names[4]string
	names[1]="mayank"
	fmt.Println(names)

	//to declare it in single line
	// nums:=[3]int{1,2,3}
	// fmt.Println(nums)

	
	// 2d arrays
	nums:=[2][2]int{{1,2},{3,4}}
	fmt.Println(nums)
}
