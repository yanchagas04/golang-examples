package main

import (
	"fmt"
	"train-go-standard/types"
)

func main() {
	fmt.Println("Hello World!")

	user := types.NewUser("John Doe", "johndoe@mail.com", "password")
	fmt.Println(user)
	user = types.NewUser("Jane Doe", "janedoe@mail.com", "password2")
	fmt.Println(user)
}
