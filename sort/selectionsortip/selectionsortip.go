package main

import "fmt"

func SelectionSortIP(v []int) {
	//criar laço para controlar a quantidade de varreduras
	for varredura := 0; varredura < len(v)-1; varredura++ {
		iMenor := varredura
		for i := varredura + 1; i < len(v); i++ {
			if v[i] < v[iMenor] {
				iMenor = i
			}
		}
		v[varredura], v[iMenor] = v[iMenor], v[varredura]
		fmt.Println("itr:", varredura, v)
	}
}

func main() {
	valores := []int{2, 8, 6, 10, 4, 5, 3}

	fmt.Println("Antes:", valores)
	SelectionSortIP(valores)
	fmt.Println("Depois:", valores)
}
