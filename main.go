package main

import (
	"fmt"

	"github.com/joaohsc/estrutura-de-dados-go/list"
)

func main() {
	call_arraylist()

}

func call_arraylist() {
	l := list.ArrayList{}
	l.Init(10)

	for i := 1; i <= 50; i++ {
		l.Add(i)
	}
	val, _ := l.Get(0)
	fmt.Println("Valor na posicao 0: ", val)

	val, _ = l.Get(49)
	fmt.Println("Valor na posicao 49: ", val)
}
