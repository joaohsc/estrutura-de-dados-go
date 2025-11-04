package main

import "fmt"

func merge(v []int, e []int, d []int) {
	indexV, indexE, indexD := 0, 0, 0
	tamE, tamD := len(e), len(d)

	for indexE < tamE && indexD < tamD {
		if e[indexE] <= d[indexD] {
			v[indexV] = e[indexE]
			indexE++
		} else {
			v[indexV] = d[indexD]
			indexD++
		}
		indexV++
	}

	for indexE < tamE {
		v[indexV] = e[indexE]
		indexE++
		indexV++
	}

	for indexD < tamD {
		v[indexV] = d[indexD]
		indexD++
		indexV++
	}
}

func mergeSort(v []int) {
	tamV := len(v)
	if tamV > 1 {
		meio := tamV / 2

		v1 := make([]int, meio)
		copy(v1, v[:meio])

		v2 := make([]int, tamV-meio)
		copy(v2, v[meio:])

		mergeSort(v1)
		mergeSort(v2)

		fmt.Println("meio:", v[meio])
		fmt.Println("esquerdo:", v1)
		fmt.Println("direito:", v2)

		merge(v, v1, v2)
	}
}

func main() {
	v := []int{7, 6, 5, 4, 3, 3, 2, 1}
	fmt.Println("Antes da ordenação:", v)
	mergeSort(v)
	fmt.Println("Depois da ordenação:", v)
}
