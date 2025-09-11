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

func (linkedlist *LinkedList) Add(value int) {
	newNode := node{value, nil}

	aux := linkedlist.head
	prev := aux

	for aux != nil {
		prev = aux
		aux = aux.prox
	}

	if prev == nil {
		linkedlist.head = &newNode
	} else {
		prev.prox = &newNode
	}
	linkedlist.size++
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

func (linkedlist *LinkedList) AddOnIndex(value int, index int) error {
	if index >= 0 && index <= linkedlist.size {
		newNode := node{value, nil}
		if index == 0 {
			if linkedlist.head != nil {
				newNode.prox = linkedlist.head
			}
			linkedlist.head = &newNode

		} else {
			prev := linkedlist.head

			for i := 0; i < index-1; i++ {
				prev = prev.prox
			}

			newNode.prox = prev.prox
			prev.prox = &newNode
		}
		linkedlist.size++
		return nil
	} else {
		if index < 0 {
			return errors.New("Can't add in linkedlist on index < 0")
		} else {
			return errors.New("Can't add in linkedlist on index > linkedlist.size")
		}
	}
}

func (linkedlist *LinkedList) RemoveOnIndex(index int) error {
	if index >= 0 && index < linkedlist.size {
		if index == 0 {
			linkedlist.head = linkedlist.head.prox
		} else {
			del := linkedlist.head
			prev := linkedlist.head
			for i := 0; i < index; i++ {
				prev = del
				del = del.prox
			}
			prev.prox = del.prox
		}

		linkedlist.size--
		return nil
	} else {
		if index < 0 {
			return errors.New("Can't remove from linkedlist on index < 0")
		} else {
			return errors.New("Can't remove from linkedlist on index >= linkedlist.size")
		}
	}
}

func (linkedlist *LinkedList) Set(value, index int) error {
	if index >= 0 && index < linkedlist.size {
		aux := linkedlist.head

		for i := 0; i < index; i++ {
			aux = aux.prox
		}
		aux.val = value
		return nil
	} else {
		if index < 0 {
			return errors.New("Can't set value in linkedlist on index < 0")
		} else {
			return errors.New("Can't set value in linkedlist on index >= linkedlist.size")
		}
	}
}

func (linkedlist *LinkedList) Size() int {
	return linkedlist.size
}

func (linkedlist *LinkedList) Reverse() {
	var prev, next *node
	aux := linkedlist.head

	for aux != nil {
		next = aux.prox
		aux.prox = prev
		prev = aux
		aux = next
	}

	linkedlist.head = prev
}
