package main

import (
	"fmt"
	"math"
)

func SelectionSortIP(v []float64) {
	for varredura := 0; varredura < len(v)-1; varredura++ {
		iMenor := varredura
		for i := varredura + 1; i < len(v); i++ {
			if v[i] < v[iMenor] {
				iMenor = i
			}

		}

		fmt.Println("menor valor:", v[iMenor])
		v[varredura], v[iMenor] = v[iMenor], v[varredura]
		fmt.Println("itr:", varredura, v)
	}
}

func main() {
	valores := []float64{3, 6, 2, 5, 4, 3, 7, 1, math.Pow10(9)}

	fmt.Println("Antes:", valores)
	SelectionSortIP(valores)
	fmt.Println("Depois:", valores)
}
