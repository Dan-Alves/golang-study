package main

import "fmt"

func main() {
	name := "Variable 1"
	var name2 string = "Variable 2"

	fmt.Println(name)
	fmt.Println(name2)

	name3, name4 := "Var3", " Var4"

	fmt.Println(name3, name4)

	const PI float32 = 3.14

	fmt.Println(PI)

}
