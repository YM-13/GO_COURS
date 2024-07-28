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

	AppendToList(root, 10)
	AppendToList(root, 30)

	AppendToList(root, 10)
	AppendToList(root, 30)

	AppendToList(root, 40)

	//PrintList(root)
	//delElemse(root)
	PrintList(root)
	fmt.Println(KthToLast(root, 2))
	KthToLast(root, 2)

}

// task1
// написать ф, которая удаляет из с. списка дубликаты
// -> 10 20 30 40
// 2 случая
// 1). можно использовать дополнительную память для буфера
// 2). нельзя использовать дополнительную память

func delElemse(root *Node) {

	for root != nil {
		//var ip *int
		ip := new(int)
		fmt.Println(*ip)
		*ip = 333
		fmt.Println(ip)
		fmt.Println(*ip)
		// fmt.Println(root)
		// fmt.Println(&root.Next)
		// fmt.Println(root.Next)

		// fmt.Println()
		// fmt.Println(root.Val)
		// fmt.Println(&root.Val)
		// fmt.Println()

		// fmt.Println()
		// fmt.Println("stop")
		v := root.Val
		//fmt.Println(v)
		//fr := &root
		prev := &root.Next // УКАЗАТЕЛЬ (АДРЕС) НА СЛЕДУЮЩИЙ УЗЕЛ
		fmt.Println(prev)
		fmt.Println(*prev)
		//fmt.Println(*prev)
		for node := root.Next; node != nil; node = node.Next {
			//fmt.Println(node.Next)
			if node.Val == v {
				prev = &node.Next
				fmt.Println("007", *prev)
			// }else {
			// 	//fr = &node
			// 	prev = &node
			}
		}
		root = root.Next
	}
}

// task2
// реализуйте алг. для нахождения в односвязном списке k-го элемента с конца
// 10->20->25->30->45
// KthTolast(root, 1) = 45
// KthTolast(root, 3) = 25

func KthToLast(root *Node, k int) int {
	node := root
	count := 0
	for node != nil {
		count += 1
		KthToLast(node, count)
		node = node.Next
	}
	fmt.Println(count)
	p_n := count - k + 1
	fmt.Println(p_n)
	
	if p_n == k {
		return node.Val
	}

	return 5
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
