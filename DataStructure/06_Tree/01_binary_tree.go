package main

import "fmt"

func main() {
	h1 := &HeroNode{id: 1, name: "宋江"}
	h2 := &HeroNode{id: 2, name: "卢俊义"}
	h3 := &HeroNode{id: 3, name: "吴用"}
	h4 := &HeroNode{id: 4, name: "公孙胜"}
	h5 := &HeroNode{id: 5, name: "关胜"}

	h1.left = h2
	h1.right = h3
	h3.right = h4
	h3.left = h5

	bt := &BinaryTree{root: h1}
	bt.preOrder()   // 1 2 3 5 4
	bt.infixOrder() // 2 1 5 3 4
	bt.postOrder()  // 2 5 4 3 1

	bt.preOrderSearch(5)
	bt.infixOrderSearch(5)
	bt.postOrderSearch(5)

	h6 := &HeroNode{id: 6, name: "林冲"}
	h5.left = h6
	bt.DeleteByNo(6)
	bt.preOrder()

}

type BinaryTree struct {
	root *HeroNode
}

func (bt *BinaryTree) preOrder() {
	fmt.Println("前序遍历")
	bt.root.preOrder()
}

func (bt *BinaryTree) infixOrder() {
	fmt.Println("中序遍历")
	bt.root.infixOrder()
}

func (bt *BinaryTree) postOrder() {
	fmt.Println("后序遍历")
	bt.root.postOrder()
}

func (bt *BinaryTree) preOrderSearch(no int) {
	heroNode := bt.root.preOrderSearch(no)
	if heroNode != nil {
		fmt.Println("前序遍历查找：id#", heroNode.id, "name#", heroNode.name)
	} else {
		fmt.Println("前序遍历查找：未找到id=", no, "的节点")
	}
}

func (bt *BinaryTree) infixOrderSearch(no int) {
	heroNode := bt.root.infixOrderSearch(no)
	if heroNode != nil {
		fmt.Println("中序遍历查找：id#", heroNode.id, "name#", heroNode.name)
	} else {
		fmt.Println("中序遍历查找：未找到id=", no, "的节点")
	}
}

func (bt *BinaryTree) postOrderSearch(no int) {
	heroNode := bt.root.postOrderSearch(no)
	if heroNode != nil {
		fmt.Println("后序遍历查找：id#", heroNode.id, "name#", heroNode.name)
	} else {
		fmt.Println("后序遍历查找：未找到id=", no, "的节点")
	}
}

func (bt *BinaryTree) DeleteByNo(no int) {
	if bt.root == nil {
		fmt.Println("空树")
	} else {
		if bt.root.id == no {
			bt.root = nil
		} else {
			bt.root.DeleteByNo(no)
		}
	}
}

type HeroNode struct {
	id    int
	name  string
	left  *HeroNode
	right *HeroNode
}

// preOrder 前序遍历 1 2 3 5 4
func (heroNode *HeroNode) preOrder() {
	fmt.Println(heroNode.id, "#", heroNode.name)
	if heroNode.left != nil {
		heroNode.left.preOrder()
	}
	if heroNode.right != nil {
		heroNode.right.preOrder()
	}
}

// infixOrder 中序遍历 2 1 5 3 4
func (heroNode *HeroNode) infixOrder() {
	if heroNode.left != nil {
		heroNode.left.infixOrder()
	}
	fmt.Println(heroNode.id, "#", heroNode.name)
	if heroNode.right != nil {
		heroNode.right.infixOrder()
	}
}

// postOrder 后序遍历 2 5 4 3 1
func (heroNode *HeroNode) postOrder() {
	if heroNode.left != nil {
		heroNode.left.postOrder()
	}
	if heroNode.right != nil {
		heroNode.right.postOrder()
	}
	fmt.Println(heroNode.id, "#", heroNode.name)
}

func (heroNode *HeroNode) preOrderSearch(no int) *HeroNode {
	if heroNode.id == no {
		return heroNode
	}
	var temp *HeroNode = nil
	if heroNode.left != nil {
		temp = heroNode.left.preOrderSearch(no)
	}

	if temp != nil {
		return temp
	}

	if heroNode.right != nil {
		temp = heroNode.right.preOrderSearch(no)
	}
	return temp
}

func (heroNode *HeroNode) infixOrderSearch(no int) *HeroNode {
	var temp *HeroNode = nil
	if heroNode.left != nil {
		temp = heroNode.left.infixOrderSearch(no)
	}
	if temp != nil {
		return temp
	}
	if heroNode.id == no {
		return heroNode
	}
	if heroNode.right != nil {
		temp = heroNode.right.infixOrderSearch(no)
	}
	return temp
}

func (heroNode *HeroNode) postOrderSearch(no int) *HeroNode {
	var temp *HeroNode = nil
	if heroNode.left != nil {
		temp = heroNode.left.postOrderSearch(no)
	}
	if temp != nil {
		return temp
	}

	if heroNode.right != nil {
		temp = heroNode.right.postOrderSearch(no)
	}

	if temp != nil {
		return temp
	}

	if heroNode.id == no {
		return heroNode
	}
	return temp
}

func (heroNode *HeroNode) DeleteByNo(no int) {
	if heroNode.left != nil && heroNode.left.id == no {
		heroNode.left = nil
		return
	}
	if heroNode.right != nil && heroNode.right.id == no {
		heroNode.right = nil
		return
	}
	if heroNode.left != nil {
		heroNode.left.DeleteByNo(no)
	}
	if heroNode.right != nil {
		heroNode.right.DeleteByNo(no)
	}
}
