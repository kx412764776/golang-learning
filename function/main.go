package main

import (
	"fmt"
)

func main() {

	sum := sum(5, 10)
	fmt.Println("sum =", sum)

	sum2, sub2 := sumAndSub(10, 10)
	fmt.Println("sum =", sum2, "difference =", sub2)

	sum3, sub3 := sumAndSub2(10, 10)
	fmt.Println("sum =", sum3, "difference =", sub3)

	res := cumulativeSum(5)
	fmt.Println("cumulative sum =", res)

	traverseString("Hello golang")
	traverseStringWithForRange("Hello golang")

	s1, s2 := 5, 10
	fmt.Printf("Before swap: s1 = %v, s2 = %v\n", s1, s2)
	Swap(&s1, &s2)
	fmt.Printf("After swap: s1 = %v, s2 = %v\n", s1, s2)

	test(1)
	test(1, 2, 3)

	funcAsType := funcAsType
	fmt.Printf("The type of funcAsType: %T\n", funcAsType)

	funcAsParams(funcAsType)

	var num myInt = 10
	fmt.Printf("Type of num: %T\n", num)

	// anonymous function
	sumAnonymous := func(num1 int, num2 int) int {
		return num1 + num2
	}

	ans := sumAnonymous(10, 20)
	fmt.Println(ans)
}

func sum(n1 int, n2 int) int {
	return n1 + n2
}

func sumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	var diff int
	if n1 >= n2 {
		diff = n1 - n2
	} else {
		diff = n2 - n1
	}
	return sum, diff
}

func sumAndSub2(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	if n1 >= n2 {
		sub = n1 - n2
	} else {
		sub = n2 - n1
	}
	return
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

func Swap(num1, num2 *int) {
	*num1, *num2 = *num2, *num1
}

// multiple parameters
func test(args ...int) {
	for i, value := range args {
		fmt.Printf("index: %d, value: %d\n", i, value)
	}
}

func funcAsType(num int) {
	fmt.Println(num)
}

func funcAsParams(params myFunc) {
	fmt.Printf("The type of param: %T\n", params)
}

type myInt int

type myFunc func(int)
