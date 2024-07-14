package main

import (
	"fmt"
)

// Linked List

// root -> node1 -> node2 -> node3 -> nil

type Node struct {
	Val int
	Next *Node // в Next передаем значение из Node
}

// 10 -> 20 -> 30

func main() {
	root := &Node{ Val: 10 } // здесь адрес составной (VAL и *NODE) структурной переменной в памяти
	AppendToList(root, 20) // передаем адрес головы листа и значение нового узла
	AppendToList(root, 30) // 1212[2323[====]6767[====]]
	AppendToList(root, 40)
	//PrintList(root)
	fmt.Println(Exists(root, 40))
	fmt.Println(Avg(root))
	fmt.Println(Length(root))
}


func PrintList(root *Node) {
	for node := root; node != nil; node = node.Next {
		fmt.Printf("%d ", node.Val)
		fmt.Printf("%d ", node.Next)
		fmt.Printf("\n")
	}

	fmt.Printf("\n")
}

func AppendToList(root *Node, val int) {
	node := root // адрес головы 0xc000052020

	// ищем последний элемент == циклу WHILE -> он будет выполняться пока будет удовлетворение условию
	for node.Next != nil {
		// fmt.Println(node.Next)
		// fmt.Println(node.Val)
		node = node.Next
		// fmt.Println(node)
	}

	node.Next = &Node{
		Val: val,
		Next: nil,
	}
	//fmt.Println(node.Next)
}


// HOME WORK
// существует ли элемент в списке?
func Exists(root *Node, needle int) bool {
	node := root // СОДЕРЖАНИЕ ГОЛОВЫ: [ ЗНАЧЕНИЕ, АДРЕС СЛЕДУЮЩЕГО ОБЪЕКТА ]
	// fmt.Println("Адрес первого объекта", &node)
	// fmt.Println("Содержание первого объекта:", node)
	// fmt.Println()
	// fmt.Println("Адрес второго объекта", &node.Next)
	// fmt.Println("Содержание второго объекта:", node.Next)

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
func Length(root *Node) int {
	node := root
	count := 0
	for node != nil {
		count += 1
		node = node.Next
	}
	return count
}