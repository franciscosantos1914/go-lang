package main

import "fmt"

func main() {
	test := [2]string{"Apple", "Orange"}
	fmt.Println(test)
	fmt.Println(test[1])
	Slices()
}

func Slices() {
	slice := []string{"Guava", "Mango", "Grapes"}
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(slice[1:3])
}
