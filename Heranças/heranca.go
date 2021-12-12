package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      uint8
	height   uint8
}

type student struct {
	person
	course     string
	university string
}

func main() {
	fmt.Println("Herança")

	p1 := person{"Jack", "Reacher", 35, 170}
	fmt.Println(p1)

	s1 := student{p1, "Computer Science", "UFRJ"}
	fmt.Println(s1)
}
