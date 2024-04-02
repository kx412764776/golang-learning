package main

import "fmt"

func main() {
	sum(5)

	traverseString("Hello golang")
}

func sum(num int) {
	var res = num
	for i := 0; i <= num; i++ {
		res += i
	}
	fmt.Println(res)
}

func traverseString(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
}
