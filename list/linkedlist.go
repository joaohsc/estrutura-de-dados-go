package list

import "errors"

type node struct {
	val  int
	prox *node
}

type LinkedList struct {
	head *node
	size int
}

func (linkedlist *LinkedList) Get(id int) (int, error) {
	if id >= 0 && id < linkedlist.size {
		aux := linkedlist.head

		for i := 0; i < id; i++ {
			aux = aux.prox
		}
		return aux.val, nil

	} else {
		if id < 0 {
			return id, errors.New("can't get value from linkedlist on index < 0")
		} else {
			return id, errors.New("can't get value from linkedlist on index >= linkedlist.size")
		}
	}
}
