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


// array
// [1, 2, 3, 4, 5] - N elems
// search = O(N)
// access (a[i]) = O(1)
// insert = O(N)
// delete = O(N)

// 1 2 3 4 5 6 7
// read a[2]


// list
// 1 -> 2 -> 3 -> 4 -> 5
// search = O(N)
// insert = O(1)
// access = O(N)
// delete = O(1)


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

// task
// написать функции

// существует ли элемент в списке?
// Exists(root *Node, needle int) bool

// среднее арифм. элементов
// Avg(root *Node) float64

// найти длину списка
// Length(root *Node) int
