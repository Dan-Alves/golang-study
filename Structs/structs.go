package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint8
}

func main() {
	var u user
	u.name = "Daniel"
	u.age = 24
	fmt.Println(u)

	addressEx := address{"Av. Rio Branco", 31}
	user2 := user{"João", 28, addressEx}
	fmt.Println(user2)

	user3 := user{name: "Fernanda"}
	fmt.Println(user3)
}
