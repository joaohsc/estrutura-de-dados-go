package main

import (
	"errors"
	"fmt"
	"math"
)

type Tree interface {
	AddNode(int)
	SearchNode(int) bool
	Min() (int, error)
	Max() (int, error)
	Height() int
	PreOrder()
	InOrder()
	PosOrder()
	//levelOrder()
	Remove(int) bool
}

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
	size int
}

func createNode(val int) *Node {
	return &Node{val: val, left: nil, right: nil}
}

func (bst *BST) AddNode(val int) {
	if bst.root == nil {
		bst.root = createNode(val)
	} else {
		bst.root.AddNode(val)
	}
	bst.size++
}

func (no *Node) AddNode(val int) {
	if val <= no.val {
		if no.left == nil {
			no.left = createNode(val)
		} else {
			no.left.AddNode(val)
		}
	} else {
		if no.right == nil {
			no.right = createNode(val)
		} else {
			no.right.AddNode(val)
		}
	}
}

func (bst *BST) SearchNode(val int) bool {
	if bst.root == nil {
		return false
	} else {
		return bst.root.SearchNode(val)
	}
}

func (no *Node) SearchNode(val int) bool {
	if val == no.val {
		return true
	} else if val < no.val {
		if no.left == nil {
			return false
		} else {
			return no.left.SearchNode(val)
		}
	} else {
		if no.right == nil {
			return false
		} else {
			return no.right.SearchNode(val)
		}
	}
}

func (bst *BST) Min() (int, error) {
	if bst.root == nil {
		return -1, errors.New("arvore estah vazia")
	} else {
		return bst.root.Min()
	}
}

func (no *Node) Min() (int, error) {
	min := no
	for min.left != nil {
		min = min.left
	}
	return min.val, nil
}

func (bst *BST) Max() (int, error) {
	if bst.root == nil {
		return -1, errors.New("arvore estah vazia")
	} else {
		max := bst.root
		for max.right != nil {
			max = max.right
		}
		return max.val, nil
	}
}

func (bst *BST) Height() int {
	if bst.root == nil {
		return 0
	} else {
		return bst.root.Height()
	}

}

func (no *Node) Height() int {

	if no.left == nil && no.right == nil {
		return 0
	}

	hRelativeLeft := 0
	if no.left != nil {
		hRelativeLeft = 1 + no.left.Height()
	}
	hRelativeRight := 0
	if no.right != nil {
		hRelativeRight = 1 + no.right.Height()
	}

	if hRelativeLeft >= hRelativeRight {
		return hRelativeLeft
	} else {
		return hRelativeRight
	}

}

func (no *Node) PreOrder() {
	fmt.Println(no.val)
	if no.left != nil {
		no.left.PreOrder()
	}
	if no.right != nil {
		no.right.PreOrder()
	}
}

func (no *Node) InOrder() {
	if no.left != nil {
		no.left.InOrder()
	}
	fmt.Println(no.val)
	if no.right != nil {
		no.right.InOrder()
	}
}

func (no *Node) PosOrder() {
	if no.left != nil {
		no.left.PosOrder()
	}
	if no.right != nil {
		no.right.PosOrder()
	}
	fmt.Println(no.val)
}

func (bst *BST) Remove(val int) error {
	if bst.root == nil {
		return errors.New("impossivel remover de arvore vazia")
	} else {
		bst.size--
		bst.root = bst.root.Remove(val)
		return nil
	}
}

func (no *Node) Remove(val int) *Node {
	if val < no.val {
		if no.left != nil {
			no.left = no.left.Remove(val)
			return no
		} else {
			return nil
		}
	} else if val > no.val {
		if no.right != nil {
			no.right = no.right.Remove(val)
			return no
		} else {
			return nil
		}
	} else {
		//achamos o no
		if no.left == nil && no.right == nil {
			//folha
			return nil
		} else if no.left != nil && no.right == nil {
			//1 filho à esquerda
			return no.left
		} else if no.left == nil && no.right != nil {
			//1 filho à direita
			return no.right
		} else {
			//no tem 2 filhos
			min, _ := no.right.Min()
			no.val = min
			no.right = no.right.Remove(min)
			return no

		}
	}
}

func (no *Node) IsBst(lower int, upper int) bool {
	if no.val < lower || no.val > upper {
		return false
	}
	if no.left != nil && !no.left.IsBst(lower, no.val) {
		return false
	}
	if no.right != nil && !no.right.IsBst(no.val, upper) {
		return false
	}
	return true
}

func main() {

	bst := &BST{}
	bst.AddNode(10)
	bst.AddNode(5)
	bst.AddNode(3)
	bst.AddNode(8)
	bst.AddNode(15)
	bst.AddNode(20)
	bst.AddNode(12)
	bst.AddNode(30)

	fmt.Println(bst.SearchNode(10))
	fmt.Println(bst.SearchNode(5))
	fmt.Println(bst.SearchNode(3))
	fmt.Println(bst.SearchNode(8))
	fmt.Println(bst.SearchNode(20))

	fmt.Println(bst.Min())
	fmt.Println(bst.Max())
	fmt.Println(bst.Height())

	fmt.Println(bst.Height())
	bst.Remove(20)
	fmt.Println(bst.SearchNode(20))
	fmt.Println(bst.Height())

	fmt.Println(bst.root.IsBst(math.MinInt, math.MaxInt))

	bst.root.left.right.val = 13

	fmt.Println(bst.root.IsBst(math.MinInt, math.MaxInt))

}
