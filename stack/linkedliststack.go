package stack

import "errors"

type node struct {
	val  int
	prox *node
}

type LinkedListStack struct {
	head *node
	size int
}

func (stack *LinkedListStack) Push(val int) {
	newNode := node{val, nil}

	aux := stack.head
	prev := aux

	for aux != nil {
		prev = aux
		aux = aux.prox
	}

	if prev == nil {
		stack.head = &newNode
	} else {
		prev.prox = &newNode
	}
	stack.size++
}

func (stack *LinkedListStack) Pop() (int, error) {
	if stack.size == 0 {
		return -1, errors.New("can't pop from empty list")
	}

	if stack.size == 1 {
		val := stack.head.val
		stack.head = nil
		stack.size--
		return val, nil
	}

	prev := stack.head
	curr := stack.head
	for curr.prox != nil {
		prev = curr
		curr = curr.prox
	}

	val := curr.val
	prev.prox = nil
	stack.size--

	return val, nil
}

func (stack *LinkedListStack) Peek() (int, error) {
	if stack.size == 0 {
		return -1, errors.New("can't pop from empty list")
	}

	if stack.size == 1 {
		val := stack.head.val
		return val, nil
	}

	curr := stack.head
	for curr.prox != nil {
		curr = curr.prox
	}

	val := curr.val

	return val, nil
}

func (stack *LinkedListStack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *LinkedListStack) Size() int {
	return stack.size
}
