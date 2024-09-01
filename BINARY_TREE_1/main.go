package main

import (
	"fmt"
)

// lc 110
// https://leetcode.com/problems/balanced-binary-tree/description/

// Given a binary tree, determine if it is
// height-balanced

// A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	diff := 0
	res := false
	for nod := root; nod != nil; nod = nod.Left {
		diff++
	}

	 for nod := root; nod != nil; nod = nod.Right {
		 diff--
	 }

	if diff < 0 {
		diff *= -1
	}

	if diff < 2 {
		res = true
	}

	return res
}

func main() {
	fmt.Println("Hello, World!")
}
