package main

import (
	"fmt"
	"unsafe"
)

const PI = 3.14

var (
	global1 string = "kx"
	global2 int    = 0
)

func main() {

	fmt.Println("global1 =", global1, "global2 =", global2)
	fmt.Println("global:", global1+fmt.Sprint(global2))

	var age int = 18
	age2 := 20

	fmt.Printf("age = %v\n", age)
	fmt.Printf("age type: %T\n", age)
	fmt.Println(age2)
	fmt.Printf("age2 type: %T\n", age2)
	fmt.Println(unsafe.Sizeof(age2))

	num := 13.45
	fmt.Printf("num = %v\n", num)
	fmt.Printf("num type: %T\n", num)

	var personName1, personAge1, personAccount1 = "Connor", 27, 3000.0
	fmt.Println("Person:", "name:", personName1, "age:", personAge1, "account:", personAccount1)

	personName2, personAge2, personAccount2 := "Stella", 22, 2500.0
	fmt.Println("Person:", "name:", personName2, "age:", personAge2, "account:", personAccount2)

	fmt.Println("constant variable: ", PI)
}
