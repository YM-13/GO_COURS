package main

import (
	"fmt"
)

func outer() {
	var inner func(i int)

	inner = func(i int) { // recursive inner function
		if i < 0 {
			return
		}
		fmt.Println(i)
		inner(i - 1)
	}

	inner(10)
}

func foo() (func() string, func(string)) {
	defer fmt.Println("the end of foo()")

	data := "hello from foo()!"

	getter := func() string {
		return data
	}

	setter := func(val string) {
		data = val
	}

	return getter, setter
}

func main1() {
	get, set := foo()
	fmt.Println(get())
	set("4242")
	fmt.Println(get())
}

type Node struct {
	Val  int
	Next *Node
}

func main2() {
	root := &Node{
		Val: 10,
		Next: &Node{
			Val: 20,
			Next: &Node{
				Val: 30,
			},
		},
	}

	_, nd := kth(root, 2)
	fmt.Println(nd.Val)
}

func kth(node *Node, k int) (int, *Node) {
	var fromEnd, fromEndNext int
	var nd *Node

	if node.Next == nil { // last element
		fromEnd = 1
	} else {
		fromEndNext, nd = kth(node.Next, k)
		fromEnd = 1 + fromEndNext
	}

	if fromEnd == k {
		return fromEnd, node
	}

	return fromEnd, nd
}

func add(a, b int) int {
	var c int
	c = a + b
	return c
}

// stackframe size ~ 32 bytes
// a
// b
// c
// return addr

// kth from last
// 1. make 2 pointers: i, j
// 2. move second pointer k items forward
// 3. while (j.Next != nil) { move i, j }
// 4. profit!

// 1 2 3 4 5 6 7 8 9 10
// i     j
//   i     j
//     i     j

type Nod struct {
	Val   int
	Left  *Nod
	Right *Nod
}

//     1
//    / \
//   2   3
//  / \
// 4   5
//    / \
//   6   7

// binary tree
// 1 - root
// c, d, f, g - leaves

// ????

func BInsert(root *Nod, val int) {
	if root != nil {
		if val < root.Val { // <
			if root.Left == nil {
				root.Left = &Nod{Val: val}
			} else {
				BInsert(root.Left, val)
			}
		} else { // >=
			if root.Right == nil {
				root.Right = &Nod{Val: val}
			} else {
				BInsert(root.Right, val)
			}
		}
	}
}

func Descr(root *Nod) string {
	if root == nil {
		return ""
	}
	return fmt.Sprintf("%d (%s) (%s)", root.Val,
		Descr(root.Left), Descr(root.Right),
	)
}

func main() {

	// fmt.Println(Descr(root))

	root2 := &Nod { Val: 5 }
	BInsert(root2, 2)
	BInsert(root2, 4)
	BInsert(root2, 10)

// 	 5
//  / \
// 2   10
//  \
// 	 4

// 1
// 	\
// 	 2
// 	  \
// 	   3
// 	    \
// 	     4


	fmt.Println(Descr(root2))

	// fmt.Println(treeSearch(root, 1))
	// fmt.Println(treeSearch(root, 5))
	// fmt.Println(treeSearch(root, 10))
}


func treeSearch(root *Nod, needle int) bool {
	if root == nil {
		return false
	}

	return (root.Val == needle) || treeSearch(root.Left, needle) ||
		treeSearch(root.Right, needle)
}

// 1. подумать и написать удобную функцию для создания б-дерева
// 2. написать функцию суммы всех элементов дерева
// 3. написать функцию, которая вернет число элементов в дереве
// 4* написать функцию, которая напечатает б-дерево.
//   a
//  /  \
// c    d


// !!!     bTree -  БИНАРНОЕ СБАЛАНСИРОВАННОЕ ДЕРЕВО ПОИСКА
// логарифмическое время

// + функция для поиска элемента в бинарном дереве поиска
// (дерево упорядочено)
