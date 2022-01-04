package main

import "fmt"

func variables() {
	var name string = "Arlinda Mayor"
	var number uint8 = 12
	const isBool bool = true
	age := 24
	fmt.Println(name)
	fmt.Println(number)
	fmt.Println(isBool)
	fmt.Print(age)

	job, salary := "Tester", 2000

	fmt.Println(job)
	fmt.Println(salary)
}

func main() {
	variables()
}
