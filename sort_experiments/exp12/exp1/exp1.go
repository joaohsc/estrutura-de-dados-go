package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(v []int) {
	n := len(v)
	for i := 0; i < n-1; i++ {
		trocou := false
		for j := 0; j < n-i-1; j++ {
			if v[j] > v[j+1] {
				temp := v[j]
				v[j] = v[j+1]
				v[j+1] = temp

				trocou = true
			}
		}

		if !trocou {
			return
		}
	}
}

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

func measureTimeQuicksort(v []int) {
	start := time.Now()
	QuickSort(v, 0, len(v)-1)
	elapsed := time.Since(start)
	fmt.Printf("Quicksort %s: %v ms", "tempo:", elapsed.Milliseconds())
}

func measureTimeMergeSort(v []int) {
	start := time.Now()
	mergeSort(v)
	elapsed := time.Since(start)
	fmt.Printf("Mergesort %s: %v ms", "tempo:", elapsed.Milliseconds())
}

func measureTimeBubble(v []int) {
	start := time.Now()
	bubbleSort(v)
	elapsed := time.Since(start)
	fmt.Printf("Bubble %s: %v ms", "tempo:", elapsed.Milliseconds())
}

func main() {
	v := []int{3, 6, 2, 5, 4, 3, 7, 8}

	measureTimeQuicksort(v)
	fmt.Printf("\n")
	measureTimeMergeSort(v)
	fmt.Printf("\n")
	measureTimeBubble(v)
}
