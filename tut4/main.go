package main

import (
	"fmt"
)

func main(){
	//fixed length array of same data data types and indexable
	//non initilaized elements are defaulted to their respective data value - here is 0 since int32
	var intArr [3]int32
	intArr[1] =123
	fmt.Println(intArr[0]) //print first element
	fmt.Println(intArr[1:3]) //print 2nd and 3rd element
	//values stored contiguous in memory
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	//immediately initialize the array using this syntax
	var intArr1 [3]int32 = [3]int32{1,2,3}
	fmt.Println(intArr1)

	// intArr2 := [...]int32{1,2,3,4,5}
}