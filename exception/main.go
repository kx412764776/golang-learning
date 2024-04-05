package main

import (
	"errors"
	"fmt"
)

func main() {
	test()
	err := test2()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("No error")
	}
	test3()
	fmt.Println("End of main")
}

// defer + recover: if an error occurs, the program will not crash
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error occurred:", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

// custom error
func test2() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		return errors.New("num2 cannot be 0")
	} else {
		res := num1 / num2
		fmt.Println(res)
	}
	return nil
}

// panic: if an error occurs, the program will crash
func test3() {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		panic("num2 cannot be 0")
	} else {
		res := num1 / num2
		fmt.Println(res)
	}
}
