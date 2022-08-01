package main

import (
	"fmt"
	"strings"
)

type Node struct {
	val  int
	next *Node
}

func (n *Node) Add(val int) {
	newNode := new(Node)
	newNode.val = val

	iter := n
	if n == nil {
		n = newNode
	} else {
		for ; iter.next != nil; iter = iter.next {
		}
		iter.next = newNode
	}
}

func (n Node) String() string {
	sb := strings.Builder{}

	for iter := &n; iter != nil; iter = iter.next {
		sb.WriteString(fmt.Sprintf("%d->", iter.val))
	}
	return sb.String()
}
