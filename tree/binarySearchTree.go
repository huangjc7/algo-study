package tree

import (
	"fmt"
)

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinarySearchTree struct {
	tree *TreeNode
}

// inOrderTraversal function will perform the traversal
func inOrderTraversal(root *TreeNode) {
	if root != nil {
		inOrderTraversal(root.Left)
		fmt.Println(root.Val)
		inOrderTraversal(root.Right)
	}
}

func (bst *BinarySearchTree) insert(data int) {
	if bst.tree == nil {
		bst.tree = &TreeNode{Val: data}
		return
	}

	p := bst.tree
	for p != nil {
		if data < p.Val {
			if p.Left == nil {
				p.Left = &TreeNode{Val: data}
				return
			}
			//否则持续向left递进
			p = p.Left
		} else if data > p.Val {
			if p.Right == nil {
				p.Right = &TreeNode{Val: data}
				return
			}
			p = p.Right
		} else {
			return
		}
	}
}

func (bst *BinarySearchTree) find(data int) bool {
	if bst.tree == nil {
		return false
	}
	p := bst.tree
	for {
		if p.Val == data {
			fmt.Println("value has been found", p.Val)
			return true
		} else if p.Val > data && p.Left != nil {
			p = p.Left
		} else if p.Val < data && p.Right != nil {
			p = p.Right
		} else {
			fmt.Println("value not found")
			return false
		}
	}
	return false
}
