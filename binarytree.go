package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(50)
	tree.Insert(200)
	tree.Insert(25)
	tree.Insert(7)
	tree.Insert(88)
	tree.Insert(150)
	fmt.Println(tree)
	fmt.Println(tree.Search(7))
}
