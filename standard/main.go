package main

import (
	"fmt"
	"log"
	"train-go-standard/types"
)

func main() {

	log.SetFlags(log.Ldate | log.Ltime)
	// fmt.Println("Hello World!")

	// // Types:
	// user := types.NewUser("John Doe", "johndoe@mail.com", "password")
	// fmt.Println(user)
	// user = types.NewUser("Jane Doe", "janedoe@mail.com", "password2")
	// fmt.Println(user)

	// // Tratamento de Erros:
	// msg, error := user.Greet("")
	// if error != nil {
	// 	log.Println(error)
	// }
	// fmt.Println(msg)

	// // Random
	// lotery := types.GenerateLotery()
	// lotery.Generate()
	// fmt.Println("Os números sorteados foram:", lotery.Results())

	// // Map sem Make (não recomendado)
	// sites := map[string]string{
	// 	"Google": "https://www.google.com",
	// 	"Bing":   "https://www.bing.com",
	// }
	// fmt.Println("Site: \tURL: ")
	// for k, v := range sites {
	// 	fmt.Printf("%s\t%s\n", k, v)
	// }
	// println(sites["Google"]) // Acessa o valor da chave Google
	// delete(sites, "Bing")    // Remove o item Bing do map
	// println(sites)
	// clear(sites) // Remove todos os itens do map
	// println(sites)

	// Pilha (Stack)
	pilha := new(types.Stack[int])
	teste, erro := pilha.Pop()
	if erro != nil {
		log.Println(erro)
	} else {
		fmt.Println(teste)
	}
	pilha.Push(1)
	pilha.Push(2)
	pilha.Push(3)
	fmt.Println(pilha)
	top, _ := pilha.Top()
	log.Println(top)
	pilha.Pop()
	top, _ = pilha.Top()
	log.Println(top)
}
