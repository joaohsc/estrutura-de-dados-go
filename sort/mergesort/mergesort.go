package main

import "fmt"

// Função merge: mescla dois slices ordenados no slice v original
func merge(v []int, e []int, d []int) {
	indexV, indexE, indexD := 0, 0, 0
	tamE, tamD := len(e), len(d)

	// Mescla os elementos comparando os dois slices
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

	// Copia o que sobrou do slice esquerdo
	for indexE < tamE {
		v[indexV] = e[indexE]
		indexE++
		indexV++
	}

	// Copia o que sobrou do slice direito
	for indexD < tamD {
		v[indexV] = d[indexD]
		indexD++
		indexV++
	}
}

// Função mergeSort: divide o slice e chama merge recursivamente
func mergeSort(v []int) {
	tamV := len(v)
	if tamV > 1 {
		meio := tamV / 2

		// Cria os slices esquerdo e direito
		v1 := make([]int, meio)
		copy(v1, v[:meio])

		v2 := make([]int, tamV-meio)
		copy(v2, v[meio:])

		// Chamadas recursivas
		mergeSort(v1)
		mergeSort(v2)

		// Mescla os slices ordenados de volta em v
		merge(v, v1, v2)
	}
}

func main() {
	v := []int{9, 3, 7, 1, 4, 6, 2, 5}
	fmt.Println("Antes da ordenação:", v)
	mergeSort(v)
	fmt.Println("Depois da ordenação:", v)
}
