package main

import (
	"fmt"
	"time"
)

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

		merge(v, v1, v2)
	}
}

func measureTimecrescemte() {
	start := time.Now()
	v := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Antes:", v)
	mergeSort(v)
	fmt.Println("Depois:", v)
	elapsed := time.Since(start)
	fmt.Printf("%s: %v ms", "tempo:", elapsed.Milliseconds())
}

func measureTimedecrescemte() {
	start := time.Now()
	v := []int{7, 6, 5, 4, 3, 3, 2, 1}
	fmt.Println("Antes:", v)
	mergeSort(v)
	fmt.Println("Depois:", v)
	elapsed := time.Since(start)
	fmt.Printf("%s: %v ms", "tempo:", elapsed.Milliseconds())
}

func main() {
	measureTimecrescemte()
	measureTimedecrescemte()
}
