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

	//PrintList(root)
	fmt.Println(Exists(root, 40))
	fmt.Println(Avg(root))
	fmt.Println(Length(root))
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
func Exists(root *Node, needle int) bool {
	node := root
	flag := false

	for node != nil {
		if node.Val == needle {
			flag = true
			break
		}
		node = node.Next
	}
	return flag
}
// среднее арифм. элементов
// Avg(root *Node) float64
func Avg(root *Node) float64 {
	node := root
	count := 0
	sum := 0

	for node != nil {
		count += 1
		sum += node.Val
		node = node.Next
	}
	return float64(sum) / float64(count)
}
// найти длину списка
// Length(root *Node) int
func Length(root *Node) int {
	node := root
	count := 0
	for node != nil {
		count += 1
		node = node.Next
	}
	return count
}
