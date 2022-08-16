package main

func reverse(n *Node) *Node {
	if n.next == nil || n == nil {
		return n
	}
	// 6 ->
	newhead := reverse(n.next)
	n.next.next = n
	n.next = nil
	return newhead
}
