/*
# Queue

Biblioteca da estrutura de dados Fila
*/
package types

import (
	"errors"
	"fmt"
)

/*
# Queue

Tipo que representa uma Fila, onde o primeiro que entra é o primeiro que sai.
*/
type Queue[T any] struct {
	items []T
}

// Cria uma nova pilha
func NewQueue[T any]() *Queue[T] {
	return new(Queue[T])
}

// Adiciona um elemento na Fila
func (q *Queue[T]) Add(item T) {
	q.items = append(q.items, item)
}

// Retira o primeiro elemento da Fila retornando esse elemento
func (q *Queue[T]) Remove() (item T, err error) {
	if q.items != nil {
		item = q.items[0]
		q.items = q.items[1:]
		return item, nil
	}
	return item, errors.New("Queue is empty")
}

// Retorna o primeiro elemento da Fila
func (q *Queue[T]) Peek() (item T, err error) {
	if q.items != nil {
		item := q.items[0]
		return item, nil
	}
	return item, errors.New("Queue is empty")
}

func (q *Queue[T]) String() string {
	if len(q.items) < 1 {
		return "Queue is empty"
	}
	var res string
	for i := range q.items {
		end_str := " -> "
		if i == len(q.items)-1 {
			end_str = ""
		}
		res += fmt.Sprint(q.items[i]) + end_str
	}
	return res
}
