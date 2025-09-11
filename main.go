package main

import (
	"fmt"

	"github.com/joaohsc/estrutura-de-dados-go/list"
	//"github.com/joaohsc/estrutura-de-dados-go/stack"
)

func main() {
	//call_arraylist()
	test_reverse_list()
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

func test_reverse_list() {
	list := list.RArrayList{}
	list.Inserted = 5
	list.Values = make([]int, list.Inserted)
	for i := 0; i < 5; i++ {
		list.Values[i] = (i + 1)
	}

	fmt.Println("Antes:", list.Values[:list.Inserted])

	list.Reverse()

	fmt.Println("Depois:", list.Values[:list.Inserted])
}
