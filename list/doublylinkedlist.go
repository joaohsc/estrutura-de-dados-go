package list

import "errors"

type DoublyLinkedList struct {
	head *Node2P
	tail *Node2P
	size int
}

type Node2P struct {
	prev  *Node2P
	value int
	next  *Node2P
}

func (doublylinkedlist *DoublyLinkedList) Add(value int) {
	newNode := Node2P{nil, value, nil}

	aux := doublylinkedlist.head
	prev := aux

	for aux != nil {
		prev = aux
		aux = aux.next
	}

	if prev == nil {
		doublylinkedlist.head = &newNode
		doublylinkedlist.tail = &newNode
	} else {
		prev.next = &newNode
		newNode.prev = prev
	}
	doublylinkedlist.size++
}

func (doublylinkedlist *DoublyLinkedList) AddOnIndex(value int, index int) error {
	if index >= 0 && index <= doublylinkedlist.size {
		newNode := Node2P{nil, value, nil}
		if index == 0 {
			if doublylinkedlist.head != nil {
				newNode.next = doublylinkedlist.head
			}
			doublylinkedlist.head = &newNode
			if doublylinkedlist.size == 0 {
				doublylinkedlist.tail = &newNode
			}
		} else {
			prev := doublylinkedlist.head
			for i := 0; i < index-1; i++ {
				prev = prev.next
			}
			newNode.next = prev.next
			newNode.prev = prev
			prev.next = &newNode
			if newNode.next != nil {
				newNode.next.prev = &newNode
			}
		}
		doublylinkedlist.size++
		return nil
	} else {
		if index < 0 {
			return errors.New("Can't add in doublylinkedlist on index < 0")
		} else {
			return errors.New("Can't add in doublylinkedlist on index > linkedlist.size")
		}
	}
}

func (doublylinkedlist *DoublyLinkedList) RemoveOnIndex(index int) error {
	if index >= 0 && index < doublylinkedlist.size {
		if index == 0 {
			doublylinkedlist.head = doublylinkedlist.head.next
			if doublylinkedlist.size == 1 {
				doublylinkedlist.tail = nil
			}
		} else {
			del := doublylinkedlist.head
			prev := doublylinkedlist.head
			for i := 0; i < index; i++ {
				prev = del
				del = del.next
			}
			prev.next = del.next
			if del.next != nil {
				del.next.prev = prev
			}
		}

		doublylinkedlist.size--
		return nil
	} else {
		if index < 0 {
			return errors.New("Can't remove from doublylinkedlist on index < 0")
		} else {
			return errors.New("Can't remove from doublylinkedlist on index >= linkedlist.size")
		}
	}
}

func (doublylinkedlist *DoublyLinkedList) Get(index int) (int, error) {
	if index >= 0 && index < doublylinkedlist.size {
		aux := doublylinkedlist.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
		return aux.value, nil
	} else {
		if index < 0 {
			return index, errors.New("Can't get value from doublylinkedlist on index < 0")
		} else {
			return index, errors.New("Can't get value from doublylinkedlist on index >= linkedlist.size")
		}
	}
}

func (doublylinkedlist *DoublyLinkedList) Set(value, index int) error {
	if index >= 0 && index < doublylinkedlist.size {
		aux := doublylinkedlist.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
		aux.value = value
		return nil
	} else {
		if index < 0 {
			return errors.New("Can't set value in doublylinkedlist on index < 0")
		} else {
			return errors.New("Can't set value in doublylinkedlist on index >= linkedlist.size")
		}
	}
}

func (doublylinkedlist *DoublyLinkedList) Size() int {
	return doublylinkedlist.size
}

func (doublylinkedlist *DoublyLinkedList) Reverse() {
	var tmp *Node2P
	atual := doublylinkedlist.head

	for atual != nil {
		tmp = atual.prev
		atual.prev = atual.next
		atual.next = tmp
		atual = atual.prev
	}

	tmp = doublylinkedlist.head
	doublylinkedlist.head = doublylinkedlist.tail
	doublylinkedlist.tail = tmp
}
