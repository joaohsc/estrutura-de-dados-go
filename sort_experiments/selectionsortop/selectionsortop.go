package main

import (
	"fmt"
	"math"
)

func SelectionSortOP(v []int) []int {
	//instanciar vetor ordenado com mesmo tamanho de v
	ordenado := make([]int, len(v))

	//criar variável MAX
	MAX := math.MaxInt

	//criar laço para controlar a quantidade de varreduras
	for varredura := 0; varredura < len(v); varredura++ {
		iMenor := 0
		for i := 1; i < len(v); i++ {
			if v[i] < v[iMenor] {
				iMenor = i
			}
		}
		ordenado[varredura] = v[iMenor]
		v[iMenor] = MAX
		fmt.Println("itr:", varredura, ordenado)
	}
	return ordenado
}

func main() {
	valores := []int{2, 8, 6, 10, 4, 5, 3}

	fmt.Println("Antes:", valores)
	NEW := SelectionSortOP(valores)
	fmt.Println("Depois:", NEW)
}
