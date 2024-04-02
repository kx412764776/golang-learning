package main

import "fmt"

func main() {

	sum := sum(5, 10)
	fmt.Println(sum)

	res := cumulativeSum(5)
	fmt.Println(res)

	traverseString("Hello golang")
	traverseStringWithForRange("Hello golang")

}

func sum(n1 int, n2 int) int {
	return n1 + n2
}

func cumulativeSum(num int) int {
	var res = num
	for i := 0; i <= num; i++ {
		res += i
	}
	return res
}

func traverseString(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
	fmt.Println()
}

func traverseStringWithForRange(s string) {
	for i, value := range s {
		fmt.Printf("%c index: %d\n", value, i)
	}
}
