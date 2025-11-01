package main

import "fmt"

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
		fmt.Println("itr: ", i, v)
		if !trocou {
			return
		}
	}
}

func main() {
	valores := []int{2, 8, 6, 10, 4, 5, 3}

	fmt.Println("Antes:", valores)
	bubbleSort(valores)
	fmt.Println("Depois:", valores)
}
