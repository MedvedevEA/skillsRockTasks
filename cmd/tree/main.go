// Добавлен метод BinarySearchTree.String() string
package main

import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BinarySearchTree struct {
	Root *TreeNode
}

// Insert вставляет новое значение в дерево
func (bst *BinarySearchTree) Insert(value int) {
	newNode := &TreeNode{Value: value}
	if bst.Root == nil {
		bst.Root = newNode
	} else {
		bst.Root.insert(newNode)
	}
}

// Вспомогательная функция вставки
func (node *TreeNode) insert(newNode *TreeNode) {
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
		} else {
			node.Left.insert(newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			node.Right.insert(newNode)
		}
	}
}

// Search ищет значение в дереве и возвращает true, если оно найдено
func (bst *BinarySearchTree) Search(value int) bool {
	return bst.Root != nil && bst.Root.search(value)
}

// Вспомогательная функция поиска
func (node *TreeNode) search(value int) bool {
	if node == nil {
		return false
	}
	if value < node.Value {
		return node.Left.search(value)
	} else if value > node.Value {
		return node.Right.search(value)
	} else {
		return true
	}
}
func (b *BinarySearchTree) String() string {
	if b.Root == nil {
		return "nil"
	}
	return b.Root.String()

}
func (n *TreeNode) String() string {
	if n == nil {
		return "nil"
	}
	if n.Left == nil && n.Right == nil {
		return fmt.Sprintf("%d", n.Value)
	}
	return fmt.Sprintf("%d (%s , %s)", n.Value, n.Left.String(), n.Right.String())

}

func (b *BinarySearchTree) DeleteRoot() {
}
func (n *TreeNode) delete() {
}

func main() {
	bst := &BinarySearchTree{}
	bst.Insert(20)
	bst.Insert(30)
	bst.Insert(50)

	bst.Insert(40)
	bst.Insert(70)
	bst.Insert(60)
	bst.Insert(80)
	fmt.Println(bst)
	bst.DeleteRoot()
	fmt.Println(bst)
	fmt.Println("Search for 40:", bst.Search(40))
	fmt.Println("Search for 100:", bst.Search(100))
}
