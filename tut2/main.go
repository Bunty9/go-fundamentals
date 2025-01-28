package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	//integer types with specific int size
	var intNum int16 = 32767
	fmt.Println(intNum)

	//float data type
	var floatNum float64 = 1234.134
	fmt.Println(floatNum)

	var floatNum2 float32 = 13.45
	//cannot add directly should covert the float 64 to float 32
	var result float32 = floatNum2 + float32(floatNum)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myString string = "Hello \nWorld"
	fmt.Println(myString)

	fmt.Println(len("A"))
	fmt.Println(len("*")) 
	fmt.Println(len("β")) //results in length 2 as beta is not in list of ascii characters and taks 2 length to store 

	//len of string with fancy characters willl require rune
	fmt.Println(utf8.RuneCountInString("β")) //results in count 1
 
	var myRune rune = 'a'  //runes require single quotes
	fmt.Println(myRune) //results in 97

	var myBool bool = false // default is false
	fmt.Println(myBool)
	
	var intNum3 int  //default is 0
	fmt.Println(intNum3)
	//default value for all ints, unsigned ints, floats and runes is 0
	//for strings is  ""
	//for booleans is false

	myVar := "text" //:= declares a variable, here the valuse is string to the data type is assumed as string with enough size to fit the value
	fmt.Println(myVar)

	var1,var2 := "text1",30
	//initialize multiple variables 
	fmt.Println(var1)
	fmt.Println(var2)

	//constants are alternative to variables , all things that are applicable to variables are applicable to constants
	//except cannot change its value once its created
	//cannot just declare a constant you also have to initialize with a value explicitly

	const pi float32 = 3.1415
	fmt.Println(pi)
}