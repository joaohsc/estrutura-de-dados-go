package main

import "fmt"

func QuickSort(v []int, ini int, fim int) {
	if ini < fim {
		fmt.Println("ini, fim", ini, fim)
		indexPivot := Partition(v, ini, fim)
		fmt.Println("indexPivot", indexPivot)
		QuickSort(v, ini, indexPivot-1)
		fmt.Println("vetor esquerdo", v)
		QuickSort(v, indexPivot+1, fim)
		fmt.Println("vetor direito", v)
	}
}

func Partition(v []int, ini int, fim int) int {
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
	v := []int{1, 2, 3, 3, 4, 5, 6, 7}

	fmt.Println("Antes:", v)
	QuickSort(v, 0, len(v)-1)
	fmt.Println("Depois:", v)
}
