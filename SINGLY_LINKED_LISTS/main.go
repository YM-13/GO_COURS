package main

import (
	"fmt"
)

// Linked List

// root -> node1 -> node2 -> node3 -> nil

type Node struct {
	Val  int
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
	root := &Node{Val: 10}
	AppendToList(root, 20)

	AppendToList(root, 10)
	AppendToList(root, 30)

	AppendToList(root, 10)
	AppendToList(root, 30)

	AppendToList(root, 40)
	AppendToList(root, 10)
	AppendToList(root, 10)
	// AppendToList(root, 10)

	PrintList(root)
	fmt.Println(KthToLast(root, 5))
	// delElemse(root)
	// PrintList(root)

	// PrintList(RemoveOne(root, 10))
	//PrintList(RemoveAll(root, 10))
}

// remove first occurence of val
// remove(10 20 30 10, 10) -> 20 30 10
func RemoveOne(root *Node, val int) *Node {
	if root.Val == val {
		return root.Next
	}

	node := root

	var prev *Node

	for {
		if node == nil {
			break
		}

		if node.Val == val {
			prev.Next = node.Next
			break
		}

		prev = node
		node = node.Next
	}

	return root
}

// remove all occurences of val
// remove(10 20 30 10, 10) -> 20 30
func RemoveAll(root *Node, val int) *Node {
	for root != nil && root.Val == val {
		root = root.Next
	}

	var prev *Node

	for node := root; node != nil; node = node.Next {
		if val == node.Val {
			prev.Next = node.Next
		}
		prev = node
	}

	return root
}

// 3 cases -
// 10 10 10 x x x x y y y
// x x x 10 10 10 y y y
// x x x y y y 10 10 10


// task1
// написать ф, которая удаляет из с. списка дубликаты
// -> 10 20 30 40
// 2 случая
// 1). можно использовать дополнительную память для буфера
// 2). нельзя использовать дополнительную память

// func delElemse(root *Node) {

// 	for root != nil {
// 		v := root.Val
// 		fmt.Println(v)
// 		prev := root
// 		for node := root.Next; node != nil; node = node.Next {
// 			//node2 := node.Next
// 			if node.Val == v {
// 				prev.Next = node.Next
// 			}else {
// 				root = root.Next
// 			}
// 			}

// 		}
// 	}
// }

// task2
// реализуйте алг. для нахождения в односвязном списке k-го элемента с конца
// 10->20->25->30->45
// KthTolast(root, 1) = 45
// KthTolast(root, 3) = 25

func KthToLast(root *Node, k int) int {
	count := 0
	for node := root; node != nil; node = node.Next {
		count += 1
	}
	sNum := count - k + 1
	for i := 1; i < sNum; i++ {
		root = root.Next
	}
	return root.Val
}

// РЕКУРСИЯ
var count int = 0
var i_c int = 0
var res int

func KthToLast2(root *Node, k int) int {

	if root != nil {
		count += 1
		KthToLast2(root.Next, k)
	} else {
		i_c = count - k + 1
		fmt.Println("ПОРЯДКОВЫЙ НОМЕР: ", i_c)
		return 0
	}

	if count == i_c {
		res = root.Val
		count -= 1
	} else {
		count -= 1
	}
	return res
}

func AppendToList(root *Node, val int) {
	node := root

	// ищем последний элемент
	for node.Next != nil {
		node = node.Next
	}

	node.Next = &Node{
		Val:  val,
		Next: nil,
	}
}

// task
// написать функции

// среднее арифм. элементов
// Avg(root *Node) float64

// найти длину списка
// Length(root *Node) int

// существует ли элемент в списке?
// Exists(root *Node, needle int) bool
//
// t = O(n) (f(n) = n)
// t1 = O(f(n))
// <=> алгоритм при n элементах
// всегда в самом худшем случае отработает быстрее, чем f(n)
// если t = O(n), то в худшем случае время работы растет линейно
// если t = O(n!), то это очень плохо(

// O(n) == O(n + 1) == O(n + C)

func Exists(node *Node, needle int) bool {
	for node != nil {
		if node.Val == needle {
			return true
		}
		node = node.Next
	}
	return false
}

func ExistsRec(node *Node, needle int) bool {
	if node == nil {
		return false
	}
	if node.Val == needle {
		return true
	}
	return ExistsRec(node.Next, needle)
}

func LengthRec(node *Node) int {
	if node == nil {
		return 0
	}
	return 1 + LengthRec(node.Next)
}
