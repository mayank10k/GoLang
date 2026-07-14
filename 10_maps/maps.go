package main

import (
	"fmt"
	"maps"
)

func main() {
	//creating map

	// m:=make(map[key]value)
	m := make(map[string]string)

	//setting an element
	m["name"] = "golang"
	m["phone"]="1234567890"

	//get an element
	fmt.Println(m["name"])

	// imp: if key doesn't exists iiin the map then it returns zeroed value

	fmt.Println(len(m))

	delete(m,"phone")
	fmt.Println(m["phone"])
	fmt.Println(len(m))

	clear(m)

	//use make fun when you dont know the element,else use this manual way(using make fun is better)
	m2:=map[string]int{"price":12,"phone":2}
	fmt.Println(m2)

	k,ok:=m2["phone"]   // k giv the value of the given key,ok return true or false
	fmt.Println(k)
	if ok{
		fmt.Println("all ok")
	}else{
		fmt.Println("not ok")
	}

	m1:=map[string]int{"price":12,"phone":2}

	fmt.Println(maps.Equal(m1,m2))
	// fmt.Println(m1==m2) not possible



}
