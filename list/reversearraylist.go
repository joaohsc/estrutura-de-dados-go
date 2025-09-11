package list

type RArrayList struct {
	Values   []int
	Inserted int
}

func (list *RArrayList) Reverse() {

	for i, j := 0, list.Inserted-1; i < j; i, j = i+1, j-1 {
		list.Values[i], list.Values[j] = list.Values[j], list.Values[i]
	}
}
