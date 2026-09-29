package types

import (
	"math/rand"
	"strconv"
)

type Lotery struct {
	chosen_number [6]int
}

func GenerateLotery() *Lotery {
	return &Lotery{}
}

func (lotery *Lotery) Generate() {
	for i := 0; i < 6; i++ {
		num := rand.Intn(100)
		lotery.chosen_number[i] = num
	}
}

func (lotery *Lotery) Results() string {
	var result string
	for i := 0; i < 6; i++ {
		result += strconv.Itoa(lotery.chosen_number[i]) + " "
	}
	return result
}
