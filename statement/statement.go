package main

import "fmt"

func main() {

	var count int = 100
	if count < 30 {
		fmt.Println("Insufficient stock")
	} else {
		fmt.Println("Sufficient stock")
	}

	if count2 := 20; count2 < 30 {
		fmt.Println("Insufficient stock")
	}

	a := 1
	switch {
	case a == 1:
		fmt.Println("aaa")
	case a == 2:
		fmt.Println("bbb")
	}

	switch b := 7; {
	case b > 6:
		fmt.Println("value > 6")
	case b <= 6:
		fmt.Println("value <= 6")
	}

	var score int = 0
	fmt.Println("Enter your score:")
	fmt.Scanln(&score)

	if score >= 90 {
		fmt.Println("Your score level: A")
	} else if score >= 80 && score <= 90 {
		fmt.Println("Your score level: B")
	} else if score >= 70 && score <= 80 {
		fmt.Println("Your score level: C")
	} else if score >= 60 && score <= 70 {
		fmt.Println("Your score level: D")
	} else if score <= 60 {
		fmt.Println("Your score level: E")
	}

	switch score / 10 {
	case 10:
		fmt.Println("Your score level: A")
	case 9:
		fmt.Println("Your score level: A")
	case 8:
		fmt.Println("Your score level: B")
	case 7:
		fmt.Println("Your score level: C")
	case 6:
		fmt.Println("Your score level: D")
	case 5:
		fmt.Println("Your score level: E")
	case 4:
		fmt.Println("Your score level: E")
	case 3:
		fmt.Println("Your score level: E")
	case 2:
		fmt.Println("Your score level: E")
	case 1:
		fmt.Println("Your score level: E")
	case 0:
		fmt.Println("Your score level: E")
		fallthrough
	default:
		fmt.Println("Input error")
	}

}
