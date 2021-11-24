package main

import "fmt"

func sum(x int8, y int8) int8 {
	return x + y
}

func mathematicalOperations(x, y int8) (int8, int8) {
	sum := x + y
	subtraction := x - y
	return sum, subtraction
}

func main() {
	x := sum(8, 9)
	fmt.Println(x)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("Text")
	fmt.Println(result)

	result1, result2 := mathematicalOperations(8, 10)
	fmt.Println(result1, result2)
}
