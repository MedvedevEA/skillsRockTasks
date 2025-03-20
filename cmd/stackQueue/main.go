// Добавлено обобщение типов данных Stack и Queue
// Добавлены методы Stake.Peek() и Queue.GetSlice()
package main

import (
	"fmt"

	"skillsRockTasks/internal/helpers"
)

type Stack[T any] struct {
	elements []T
}

// Конструктор Stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		elements: []T{},
	}
}

// Push добавляет элемент в стек
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop удаляет верхний элемент из стека и возвращает его
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		return helpers.GetZero[T](), false
	}
	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]
	return element, true
}

// Peek получить первый элемент стека без удаления
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) > 0 {
		return s.elements[0], true
	}
	return helpers.GetZero[T](), false
}

type Queue[T any] struct {
	elements []T
}

// Конструктор Queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		elements: []T{},
	}
}

// Enqueue добавляет элемент в конец очереди
func (q *Queue[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
}

// Dequeue удаляет элемент из начала очереди и возвращает его
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.elements) == 0 {
		return helpers.GetZero[T](), false
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, true
}

// GetSlice
func (q *Queue[T]) GetSlice() []T {
	return q.elements
}

func main() {
	stack := NewStack[string]()
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	fmt.Println(stack.Peek()) // Вывод: a, true
	fmt.Println(stack.Pop())  // Вывод: c, true
	fmt.Println(stack.Pop())  // Вывод: b, true
	fmt.Println(stack.Pop())  // Вывод: a, true
	fmt.Println(stack.Pop())  // Вывод: , false

	queue := NewQueue[bool]()
	queue.Enqueue(true)
	queue.Enqueue(false)
	queue.Enqueue(true)
	fmt.Println(queue.GetSlice()) // Вывод: true, false, true
	fmt.Println(queue.Dequeue())  // Вывод: true, true
	fmt.Println(queue.Dequeue())  // Вывод: false, true
	fmt.Println(queue.Dequeue())  // Вывод: true, true
	fmt.Println(queue.Dequeue())  // Вывод: false, false
}
