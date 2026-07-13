package main

import (
	"fmt"
	"slices"
)

//most used construct in go


func main(){
	//declaration

	//uninitialized slice in nil (null)
	// var nums[]int
	// fmt.Println(nums)
	// fmt.Println(nums==nil)
	// fmt.Println(len(nums))

    // make is inbuild function to initialise slice in go 
	// var arr=make([]int,2)   //make(datatype,initial size)
	// fmt.Println(arr==nil)

	var arr =make([]int,2,5);    // make(datatype,initial size ,initial capacity)
	fmt.Println("Length: ",len(arr))
	fmt.Println("Capacity: ",cap(arr))

	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	fmt.Println(cap(arr))

	arr = append(arr, 4)
	fmt.Println(cap(arr))


	//copy functioin

	// var nums=make([]int,0,5)
	// nums=append(nums, 2)

	// var nums2=make([]int,len(nums))
	// copy(nums2,nums)
	
	// fmt.Println(nums,nums2)


	// slice operator

	var num=[]int {1,2,3}

	fmt.Println(num[0:2])  //num[x:y] from x index to y-1 index,(x is excluded)
	fmt.Println(num[:2])   // from start to 2-1,(2 index is excluded)
	fmt.Println(num[1:])    // from 1 to last


	// slice package

	var nums1=[]int{1,2}
	var nums2=[]int{1,2}

	fmt.Println(slices.Equal(nums1,nums2))
	// fmt.Println(nums1==nums2)   these can not be compared like this ,slice can only be compared with 'nil'


	// 2D slices

	var nums=[][]int{{1,2,3},{4,5,6}}
	fmt.Println(nums)

}
