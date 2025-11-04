package main

import (
	"fmt"
	"math"
)

func InsertionSortIP(v []float64) {
	for i := 1; i < len(v); i++ {
		for j := i; j > 0; j-- {
			if v[j-1] > v[j] {
				fmt.Println("i:", i)
				fmt.Println("troca:", v[j-1], v[j])
				v[j-1], v[j] = v[j], v[j-1]
			} else {
				break
			}
		}
		fmt.Println("itr:", i, v)
	}
}

func main() {
	v := []float64{3, 6, 2, 5, 4, 3, 7, 1, math.Pow10(9)}
	fmt.Println("Original:", v)
	InsertionSortIP(v)
	fmt.Println("Ordenado:", v)
}
