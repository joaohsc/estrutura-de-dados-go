package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(v []int, ini int, fim int) {
	if ini < fim {
		indexPivot := Partition(v, ini, fim)
		QuickSort(v, ini, indexPivot-1)
		QuickSort(v, indexPivot+1, fim)
	}
}

func Partition(v []int, ini int, fim int) int {
	randIndex := ini + rand.Intn(fim-ini+1)
	v[randIndex], v[fim] = v[fim], v[randIndex]

	pivot := v[fim]
	pIndex := ini

	for i := ini; i < fim; i++ {
		if v[i] <= pivot {
			v[pIndex], v[i] = v[i], v[pIndex]
			pIndex++
		}
	}

	v[pIndex], v[fim] = v[fim], v[pIndex]
	return pIndex
}

func main() {
	v := []int{9, 3, 7, 1, 5, 2, 8, 4, 6}

	fmt.Println("Antes:", v)
	QuickSort(v, 0, len(v)-1)
	fmt.Println("Depois:", v)
}
