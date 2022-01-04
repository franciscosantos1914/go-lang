package main

import "fmt"

func main() {
	// IfAndElse()
	// Switches("yellow")
	Loops()
}

func Loops() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number %d\n", i)
	}
}

func Switches(color string) {
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not red or blue")
	}
}

func IfAndElse() {
	x := 15
	y := 10

	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}

	color := "yellow"

	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is not red or blue")
	}

}
