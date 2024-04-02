package main

import (
	"fmt"
	"strconv"
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

	// char
	var char byte = 'a'
	fmt.Println(char) //97
	fmt.Printf("char: %c\n", char)

	// type casting
	var n int = 100
	fmt.Printf("type: %T\n", n)
	cast := float64(n)
	fmt.Printf("type after casting: %T\n", cast)

	var toString1 int = 10
	var toString2 float32 = 12.34
	var toString3 bool = false
	var toString4 byte = 'x'

	var s1 = fmt.Sprintf("%d", toString1)
	var s2 = fmt.Sprintf("%f", toString2)
	var s3 = fmt.Sprintf("%t", toString3)
	var s4 = fmt.Sprintf("%c", toString4)
	fmt.Printf("s1: %q, type: %T\n", s1, s1)
	fmt.Printf("s2: %q, type: %T\n", s2, s2)
	fmt.Printf("s3: %q, type: %T\n", s3, s3)
	fmt.Printf("s4: %q, type: %T\n", s4, s4)

	// string -> int
	var intString = "19"
	var stringToInt, _ = strconv.ParseInt(intString, 10, 0)
	fmt.Printf("value: %v, type: %T\n", stringToInt, stringToInt)
}
