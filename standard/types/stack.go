package types

import (
	"errors"
	"fmt"
)

/*
# Interface

As interfaces em Go, assim como em outras linguagens, servem para verificar se seu tipo (ou "Classe") implementa todas as funções definidas por essa interface.
  - Funciona como um contrato que o tipo deve seguir para ser considerado do mesmo tipo da interface.
  - Permite que você utilize vários tipos que implementam as funções dessa interface como se fossem do mesmo tipo.
*/
type StackInterface[T any] interface {
	Top() (T, error)
	Pop() (T, error)
	Push(item T)
}

/*
# Stack

Tipo que representa uma Pilha, onde o último item que entra é o primeiro que sai.
*/
type Stack[T any] struct {
	items []T
}

// Cria um nova pilha
func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

// Retira o último elemento da pilha
func (s *Stack[T]) Pop() (item T, err error) {
	if s.items != nil {
		item = s.items[len(s.items)-1]     // Pega o último item
		s.items = s.items[:len(s.items)-1] // Remove ele da stack
		return item, nil
	}
	return item, errors.New("Stack is empty")
}

// Adiciona um elemento na pilha
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Retorna o último elemento da pilha
func (s *Stack[T]) Top() (item T, err error) {
	if s.items != nil {
		return s.items[len(s.items)-1], nil
	}
	return item, errors.New("Stack is empty")
}

func (s *Stack[T]) String() string {
	if len(s.items) < 1 {
		return "Stack is empty"
	}
	var res string
	for i := range s.items {
		res += "\t" + fmt.Sprint(s.items[i]) + "\n"
	}
	return res
}
