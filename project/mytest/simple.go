package main

import "fmt"

func isEven(x int) bool {
	if x%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	if isEven(2) {
		fmt.Println("It is even")
	} else {
		fmt.Println("it is odd")
	}
}
