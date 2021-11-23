package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing file main")
	auxiliar.Write()

	err := checkmail.ValidateFormat("danielalves.ccomp")

	fmt.Println(err)
}
