package main

import (
	"fmt"
)

func main() {
	greet := greeting("Arly")
	fmt.Println(greet)
	getSum(2, 4, 6)
}

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2, num3 int16) {
	sum := num1 + num2 + num3
	fmt.Println(sum)
}
