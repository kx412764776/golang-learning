package main

import "fmt"

func main() {
	var age int = 18
	fmt.Println("age address:", &age)

	// var pointer *int = age -> error: cannot use age (variable of type int) as *int value in variable declaration
	// var pointer *float32 = &age -> cannot use &age (value of type *int) as *float32 value in variable declaration
	var pointer *int = &age
	fmt.Println("Address pointed to by pointer:", pointer)
	fmt.Println("pointer address:", &pointer)
	fmt.Println("Value pointed to by pointer:", *pointer)

	*pointer = 20
	fmt.Println(age)

}
