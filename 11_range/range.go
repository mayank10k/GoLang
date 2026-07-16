package main

import "fmt"

//used for iterating over data sstructures
func main() {
	nums:=[]int{1,2,3,4}

	// for i:=0; i<len(nums);i++{
	// 	fmt.Println(nums[i])
	// }

	sum:=0

	for i,num:=range nums{  //(index,val)--> if used only once it will give the index value
		sum+=num
		fmt.Println(i,num)
	}

	fmt.Println(sum)

	m:=map[string]string{"first":"mayank","last":"mehta"}

	for k,v:= range m{      //if only one var is used it give the key only
		fmt.Println(k,v)
	}


	//how range is used in strings

	// i is the starting byte of rune

	for i,c :=range "golang"{    // c will give3 the unicode code point rune
		// fmt.Println(i,c)
		fmt.Println(i,string(c))
	}

	// c:='a'
	var c rune='a'
	fmt.Println(c)
	fmt.Printf("%c\n",c)
	 


}
