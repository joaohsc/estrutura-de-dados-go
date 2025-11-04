package main

import "fmt"

func CountingSort(v []int) []int {
	if len(v) == 0 {
		return v
	}

	max := v[0]
	for _, val := range v {
		if val > max {
			max = val
		}
	}

	count := make([]int, max+1)

	for _, val := range v {
		count[val]++
	}

	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	output := make([]int, len(v))
	for i := len(v) - 1; i >= 0; i-- {
		val := v[i]
		output[count[val]-1] = val
		count[val]--
	}

	return output
}

func main() {
	v := []int{1, 2, 3, 3, 4, 5, 6, 7}
	sorted := CountingSort(v)
	fmt.Println("Vetor original:", v)
	fmt.Println("Vetor ordenado:", sorted)
}
