package main

import (
	"errors"
	"fmt"
)

func main() {

	var number0 int8 = 100
	var number1 int16 = 10000
	var number2 int32 = 1000000000
	var number3 int64 = 1000000000000000000
	var number4 int = 1000000000000000000

	fmt.Println(number0, number1, number2, number3, number4)

	var number5 float32 = 123.5
	var number6 float64 = 123456789.9

	fmt.Println(number5, number6)

	var word string = "Test"
	fmt.Println(word)

	char := 'c'
	fmt.Println(char)

	var boolean1 bool
	fmt.Println(boolean1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
