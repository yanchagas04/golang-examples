package main

import (
	"fmt"
	"log"
	"train-go-standard/types"
)

func main() {
	log.SetPrefix("train-go: ")
	log.SetFlags(0)

	fmt.Println("Hello World!")

	// Types:
	user := types.NewUser("John Doe", "johndoe@mail.com", "password")
	fmt.Println(user)
	user = types.NewUser("Jane Doe", "janedoe@mail.com", "password2")
	fmt.Println(user)

	// Tratamento de Erros:
	msg, error := user.Greet("")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(msg)
}
