package main

import (
	"fmt"
	"math"
)

func bubbleSort(v []float64) {
	n := len(v)
	for i := 0; i < n-1; i++ {
		trocou := false
		for j := 0; j < n-i-1; j++ {
			if v[j] > v[j+1] {
				temp := v[j]
				v[j] = v[j+1]
				v[j+1] = temp

				trocou = true
				if trocou {
					fmt.Println("trocou? ", trocou)
					fmt.Println("troca: ", temp, v[j])
				}
			}
		}

		fmt.Println("itr: ", i, v)
		if !trocou {
			return
		}
	}
}

func main() {
	valores := []float64{3, 6, 2, 5, 4, 3, 7, 1, math.Pow10(9)}

	fmt.Println("Antes:", valores)
	bubbleSort(valores)
	fmt.Println("Depois:", valores)
}
