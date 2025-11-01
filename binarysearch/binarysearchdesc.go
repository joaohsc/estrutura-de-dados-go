package binarysearch

func revbinsearch(val int, list []int) int {
	start := 0
	end := list.Size() - 1

	for start <= end {
		// Encontra o meio
		mid := start + (end-start)/2

		// Pega o valor no meio
		midVal := list.Get(mid)

		if midVal == val {
			// Valor encontrado, retorna o índice
			return mid
		} else if midVal < val {
			// O valor procurado é maior, então ele deve estar na primeira metade da lista
			// (onde estão os valores maiores)
			end = mid - 1
		} else { // midVal > val
			// O valor procurado é menor, então ele deve estar na segunda metade da lista
			// (onde estão os valores menores)
			start = mid + 1
		}
	}

	// Valor não encontrado
	return -1
}
