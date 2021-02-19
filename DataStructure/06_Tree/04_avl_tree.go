package main

import "fmt"

func main() {

	//arr := []int{4, 3, 6, 5, 7, 8} // 左旋
	//arr := []int{10, 12, 8, 9, 7, 6} // 右旋
	arr := []int{10, 11, 7, 6, 8, 9} // 双旋
	tree := &AVLTree{}

	for i := 0; i < len(arr); i++ {
		tree.Add(arr[i])
	}

	tree.InfixOrder()

	fmt.Println("树的高度#", tree.Height())
	fmt.Println("左子树高度#", tree.leftHeight())
	fmt.Println("右子树高度#", tree.rightHeight())

	fmt.Println("当前根节点#", tree.root.id)
	fmt.Println("当前根节点的左子节点#", tree.root.left.id)
	fmt.Println("当前根节点的右子节点#", tree.root.right.id)

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

	// 当添加完一个节点后，如果刚加入的节点的右节点的高度与左节点的高度差大于1的话，向左旋转
	if node.rightHeight()-node.leftHeight() > 1 {
		// 如果当前节点的右子节点的左子节点高度大于当前节点的右子节点的右子节点的高度
		if node.right != nil && node.right.leftHeight() > node.right.rightHeight() {
			// 对当前节点的右节点进行向右旋转
			node.right.rightRotate()
		}
		node.leftRotate()
		return
	}

	// 当添加完一个节点后，如果刚加入的节点的左节点的高度与右节点的高度差大于1的话，向右旋转
	if node.leftHeight()-node.rightHeight() > 1 {
		// 如果当前节点的左子节点的右子节点高度大于当前节点的左子节点的左子节点的高度
		if node.left != nil && node.left.rightHeight() > node.left.leftHeight() {
			// 对当前节点的左节点进行向左旋转
			node.left.leftRotate()
		}
		node.rightRotate()
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

func (node *AVLNode) leftRotate() {
	// 以当前节点的值，创建新的节点。
	newNode := &AVLNode{id: node.id}
	// 把新的节点的左子节点设置为当前节点的左子节点
	newNode.left = node.left
	// 把新的节点的右子节点设置为当前节点的右子节点的左子节点
	newNode.right = node.right.left
	// 把当前节点的值替换成右子节点的值
	node.id = node.right.id
	// 把当前节点的右子节点设置为当前节点的右子节点的右子节点
	node.right = node.right.right
	// 把当前节点的左子节点设置为新的节点
	node.left = newNode
}

func (node *AVLNode) rightRotate() {
	// 以当前节点的值，创建新的节点。
	newNode := &AVLNode{id: node.id}
	// 把新的节点的右子节点设置为当前节点的右子节点
	newNode.right = node.right
	// 把新的节点的左子节点设置为当前节点的左子节点的右子节点
	newNode.left = node.left.right
	// 把当前节点的值替换成左子节点的值
	node.id = node.left.id
	// 把当前节点的左子节点设置为当前节点的左子节点的左子节点
	node.left = node.left.left
	// 把当前节点的右子节点设置为新的节点
	node.right = newNode
}
