package main

import "fmt"

func main() {

	arr := []int{4, 3, 6, 5, 7, 8}
	tree := &AVLTree{}

	for i := 0; i < len(arr); i++ {
		tree.Add(arr[i])
	}

	tree.InfixOrder()

	fmt.Println("树的高度#", tree.Height())
	fmt.Println("左子树高度#", tree.leftHeight())
	fmt.Println("右子树高度#", tree.rightHeight())

}

type AVLTree struct {
	root *AVLNode
}

func (tree *AVLTree) Add(id int) {
	if tree.root == nil {
		tree.root = &AVLNode{id: id}
	} else {
		tree.root.Add(id)
	}
}

func (tree *AVLTree) InfixOrder() {
	tree.root.infixOrder()
}

func (tree *AVLTree) Height() int {
	return tree.root.Height()
}

func (tree *AVLTree) leftHeight() int {
	if tree.root.left == nil {
		return 0
	}
	return tree.root.leftHeight()
}

func (tree *AVLTree) rightHeight() int {
	if tree.root.right == nil {
		return 0
	}
	return tree.root.rightHeight()
}

type AVLNode struct {
	id    int
	left  *AVLNode
	right *AVLNode
}

func (node *AVLNode) leftHeight() int {
	if node.left == nil {
		return 0
	}
	return node.left.Height()
}

func (node *AVLNode) rightHeight() int {
	if node.right == nil {
		return 0
	}
	return node.right.Height()
}

// Height 以当前节点为根节点的树的高度
func (node *AVLNode) Height() int {
	var leftHeight int
	if node.left == nil {
		leftHeight = 0
	} else {
		leftHeight = node.left.Height()
	}

	var rightHeight int
	if node.right == nil {
		rightHeight = 0
	} else {
		rightHeight = node.right.Height()
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func (node *AVLNode) Add(id int) {
	n := &AVLNode{id: id}

	if id < node.id {
		if node.left == nil {
			node.left = n
		} else {
			node.left.Add(id)
		}
	} else {
		if node.right == nil {
			node.right = n
		} else {
			node.right.Add(id)
		}
	}

}

func (node *AVLNode) infixOrder() {
	if node.left != nil {
		node.left.infixOrder()
	}
	fmt.Println("node id#", node.id)
	if node.right != nil {
		node.right.infixOrder()
	}
}
