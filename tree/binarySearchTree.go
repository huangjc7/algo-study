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

func preOrderTraversal(root *TreeNode) {
	if root != nil {
		fmt.Println(root.Val)
		preOrderTraversal(root.Left)
		preOrderTraversal(root.Right)
	}
}

func postOrderTraversal(root *TreeNode) {
	if root != nil {
		postOrderTraversal(root.Left)
		postOrderTraversal(root.Right)
		fmt.Println(root.Val)
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
	p := bst.tree
	for p != nil {
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

func (bst *BinarySearchTree) delete(data int) {
	// record current node
	p := bst.tree
	// record current node father
	var pp *TreeNode

	//      4
	//     / \
	//    2   6
	//   / \ /
	//  1  3 5

	for p != nil {
		// judge leaf node
		if data < p.Val {
			pp = p
			p = p.Left

		} else if data > p.Val {
			pp = p
			p = p.Right
			// must is leaf node

		} else {
			if p.Left != nil && p.Right != nil { // Node has two children
				// Find the minimum node on the right subtree
				minP := p.Right
				minPP := p // parent of minP
				for minP.Left != nil {
					minPP = minP
					minP = minP.Left
				}
				p.Val = minP.Val // Replace value
				p = minP         // Move to minP to delete it
				pp = minPP       // Let the parent node pointer of the initially deleted element point to the new node
			}

			//node 0 or 1 child
			var child *TreeNode
			if p.Left != nil {
				child = p.Left
			} else if p.Right != nil {
				child = p.Right
			}

			// 判断p的pp是不是nil 如果是nil就是root节点
			if pp == nil {
				bst.tree = child
			} else if pp.Left == p {
				pp.Left = child
			} else {
				pp.Right = child
			}
			return
		}
	}
	fmt.Println("value not found")
}
