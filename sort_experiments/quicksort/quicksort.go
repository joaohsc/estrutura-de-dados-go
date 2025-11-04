package main

import (
	"fmt"
	"math/rand"
	"time"
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

func measureTime() {
	start := time.Now()
	v := []int{1, 2, 3, 3, 4, 5, 6, 7}
	fmt.Println("Antes:", v)
	QuickSort(v, 0, len(v)-1)
	fmt.Println("Depois:", v)
	elapsed := time.Since(start)
	fmt.Printf("%s: %v ms", "tempo:", elapsed.Milliseconds())
}

func main() {
	measureTime()
}
