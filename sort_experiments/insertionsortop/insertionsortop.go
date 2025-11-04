package main

import (
	"fmt"
)

func InsertionSortOP(v []int) []int {
	if len(v) == 0 {
		return []int{}
	}

	// Criar o vetor ordenado
	ordenado := make([]int, len(v))
	ordenado[0] = v[0] // primeiro elemento já está "ordenado"

	for i := 1; i < len(v); i++ { // i = índice do elemento a inserir
		j := 0
		// encontrar a posição correta na parte ordenada
		for j = 0; j < i; j++ {
			if ordenado[j] > v[i] { // encontramos a posição da inserção
				// abrir espaço deslocando elementos à direita
				for k := i; k > j; k-- {
					ordenado[k] = ordenado[k-1]
				}
				break
			}
		}
		ordenado[j] = v[i] // inserir o elemento na posição correta
	}

	return ordenado
}

func main() {
	v := []int{3, 5, 5, 4}
	fmt.Println("Original:", v)
	sorted := InsertionSortOP(v)
	fmt.Println("Ordenado:", sorted)
}
