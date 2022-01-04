package main

import (
	"fmt"
	"strconv"
)

func main() {
	// sum := Adders()
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(sum(i))
	// }
	person1 := Person{name: "Francisco", age: 24, city: "Luanda", gender: "M"}
	fmt.Println(person1.greet())
}

func Map() {
	emails := make(map[string]string)
	emails["Arly"] = "arly@santos.com"
	emails["Francisco"] = "francisco@santos.com"

	name := map[string]string{"Arly": "arly@santos.com", "Francisco": "francisco@santos.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Arly"])
	fmt.Println(name)

}

func Ranges() {
	ids := []int{33, 76, 54, 23, 11, 2}
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
}

func Pointers() {
	a := 5
	b := &a
	*b = 10
	fmt.Println(a, *b)
}

func Adders() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

type Person struct {
	name   string
	age    int
	city   string
	gender string
}

func (p Person) greet() string {
	return "Hello, my name is " + p.name + " and I am " + strconv.Itoa(p.age) + " years old"
}
func (p *Person) hasBirthday() {
	p.age++
}
