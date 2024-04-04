package main

import "fmt"

func main() {
	str := "Hello, Golang"
	fmt.Println(len(str))

	for i, value := range str {
		fmt.Printf("index: %d, value: %c\n", i, value)
	}

	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i])
	}
}
