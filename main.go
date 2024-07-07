package main

import (
	"fmt"
)

// Linked List

// root -> node1 -> node2 -> node3 -> nil

type Node struct {
	Val int
	Next *Node
}

// 10 -> 20 -> 30





func main() {
	root := &Node{ Val: 10 }
	AppendToList(root, 20)
	AppendToList(root, 30)
	AppendToList(root, 40)

	PrintList(root)
}


func PrintList(root *Node) {
	for node := root; node != nil; node = node.Next {
		fmt.Printf("%d ", node.Val)
	}

	fmt.Printf("\n")
}

func AppendToList(root *Node, val int) {
	node := root

	// ищем последний элемент
	for node.Next != nil {
		node = node.Next
	}

	node.Next = &Node{
		Val: val,
		Next: nil,
	}
}
