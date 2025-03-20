// Добавлен метод LinkedList.Delete() bool, Invert()
// Добавлен тип Linked2List
// Добавлен метод Linked2List.Prerend(value), Linked2List.Print()
package main

import (
	"fmt"
)

type Node2 struct {
	Value    int
	Next     *Node2
	Previous *Node2
}
type Linked2List struct {
	Head *Node2
}

// Добавление элемента в начало списка
func (list *Linked2List) Prepend(value int) {
	newNode2 := &Node2{Value: value, Next: list.Head, Previous: nil}
	if list.Head != nil {
		list.Head.Previous = newNode2
	}
	list.Head = newNode2
}

// Печать всех элементов списка
func (list *Linked2List) Print() {
	fmt.Print("nil")
	for node := list.Head; node != nil; node = node.Next {
		fmt.Printf(" - %d", node.Value)

	}
	fmt.Println(" - nil")
}

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

// Добавление элемента в начало списка
func (list *LinkedList) Prepend(value int) {
	newNode := &Node{Value: value, Next: list.Head}
	list.Head = newNode
}

// Печать всех элементов списка
func (list *LinkedList) Print() {
	for node := list.Head; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Value)
	}
	fmt.Println("nil")
}

// Инвертирования всех элементов списка
func (list *LinkedList) Invert() {
	var previousNode *Node = nil
	var node = list.Head
	for node != nil {
		node, node.Next, previousNode = node.Next, previousNode, node
	}
	list.Head = previousNode
}

// Удаление элемента
func (list *LinkedList) Delete() bool {
	if list.Head == nil {
		return false
	}
	list.Head = list.Head.Next
	return true

}
func main() {
	fmt.Println("* * * LinkedList * * *")
	list := &LinkedList{}
	list.Prepend(1)
	list.Prepend(2)
	list.Prepend(3)
	list.Print() // Вывод: 3 -> 2 -> 1 -> nil
	list.Invert()
	list.Print() // Вывод: 1 -> 2 -> 3 -> nil
	list.Delete()
	list.Print() // Вывод: 2 -> 3 -> nil
	fmt.Println("* * * Linked2List * * *")
	list2 := &Linked2List{}
	list2.Print()
	list2.Prepend(1)
	list2.Prepend(2)
	list2.Prepend(3)
	list2.Print() //nil - 3 - 2 - 1 - nil
}
