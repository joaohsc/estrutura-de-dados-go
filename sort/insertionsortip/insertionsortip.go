package main

import (
	"fmt"
)

func InsertionSortIP(v []int) {
	for i := 1; i < len(v); i++ {
		for j := i; j > 0; j-- {
			if v[j-1] > v[j] {
				v[j-1], v[j] = v[j], v[j-1]
			} else {
				break
			}
		}
	}
}

func main() {
	v := []int{3, 5, 5, 4}
	fmt.Println("Original:", v)
	InsertionSortIP(v)
	fmt.Println("Ordenado:", v)
}
