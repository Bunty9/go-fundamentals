package main

import (
	"errors"
	"fmt"
)

func main(){
	printValue := "hello world"
	printMe(printValue)
	numerator :=10
	denominator :=2
	var result, remainder,err = intDivision(numerator, denominator)
	// if err!=nil{
	// 	fmt.Println(err.Error())
	// }else if remainder ==0{
	// 	fmt.Printf("The result of the division is %v", result)
	// }else{
	// 		fmt.Printf("the result of the division is %v with remainer %v", result,remainder)
	// }

	switch  {
	case err!=nil:
		fmt.Println(err.Error())
	case remainder ==0:
		fmt.Printf("the result of the division is %v", result)
	default:
		fmt.Printf("the result of the division is %v with remainder %v ", result,remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("the division was exact")
	case 1,2:
		fmt.Println("the division was close")
	default:
		fmt.Println("the division was not close")
	}

	//can pass multiple checks in if statement.
	//go only checks the min conditions and ignores the rest if passed
	if 1==1 || 2==2{ //here in or statement only checks 1==1 as it is true and can proceed with the code, the 2==2 condition check is ignored/skipped
		fmt.Println("passed the check")
	} 
}

func printMe(printvalue string){
	fmt.Println(printvalue)
}

func intDivision(numerator int, denominator int) (int,int,error){
	var err error
	if denominator==0{
		err = errors.New("cannot divide by 0")
		return 0,0,err
	}
	result := numerator/denominator
	remainder := numerator%denominator
	return result,remainder,err
	
}