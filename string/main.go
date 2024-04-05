package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "Hello, Golang"
	fmt.Println(len(str))

	for i, value := range str {
		fmt.Printf("index: %d, value: %c\n", i, value)
	}

	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\n", r[i])
	}

	// string -> int
	fmt.Println("---------string -> int---------")
	num, _ := strconv.Atoi("66")
	fmt.Printf("%v type: %T\n", num, num)

	// int -> string
	fmt.Println("---------int -> string---------")
	s := strconv.Itoa(66)
	fmt.Printf("%s type: %T\n", s, s)

	// strings
	fmt.Println("---------strings---------")
	flag1 := strings.Contains("java and golang", "go")
	fmt.Println(flag1)

	count := strings.Count("go, golang, java", "go")
	fmt.Println(count)

	// compare strings case-insensitive
	flag2 := strings.EqualFold("Go", "go")
	fmt.Println(flag2)

	// compare strings case-sensitive
	flag3 := strings.Compare("Go", "go")
	fmt.Println(flag3)
	flag4 := strings.Compare("go", "go")
	fmt.Println(flag4)

	// first index of substring
	index := strings.Index("go, golang, java", "go")
	fmt.Println(index)

	// last index of substring
	lastIndex := strings.LastIndex("go, golang, java", "go")
	fmt.Println(lastIndex)

	// replace
	newStr := strings.Replace("go, golang, java", "go", "python", -1)
	fmt.Println(newStr)

	// split
	split := strings.Split("go, golang, java", ", ")
	fmt.Println(split)

	// Lowercase, Uppercase
	lower := strings.ToLower("Go")
	fmt.Println(lower)
	upper := strings.ToUpper("Go")
	fmt.Println(upper)

	// Trim
	trim := strings.Trim("   Go   ", " ")
	fmt.Println(trim)
	trim2 := strings.Trim("~ Go ~", "~")
	fmt.Println(trim2)

	// TrimLeft
	trimLeft := strings.TrimLeft("   Go   ", " ")
	fmt.Println(trimLeft)

	// TrimSpace
	trimSpace := strings.TrimSpace("   Go and Golang ")
	fmt.Println(trimSpace)

	// HasPrefix
	hasPrefix := strings.HasPrefix("Go and Golang", "Go")
	fmt.Println(hasPrefix)

	// HasSuffix
	hasSuffix := strings.HasSuffix("Go and Golang", "Golang")
	fmt.Println(hasSuffix)
}
