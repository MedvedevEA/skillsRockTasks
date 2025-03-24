// Добавлена функция NewHashTable(int)
// Добавлены методы HashTable.Size() int, HashTable.Resize(int) *HashTable
package main

import (
	"fmt"
)

// HashTable структура для хеш-таблицы
type HashTable struct {
	array []*bucket
}

// bucket структура, представляющая список для разрешения коллизий
type bucket struct {
	head *bucketNode
}

// bucketNode структура узла в списке
type bucketNode struct {
	key   string
	value int
	next  *bucketNode
}

func NewHashTable(size int) *HashTable {
	hashTable := &HashTable{
		array: make([]*bucket, size),
	}
	for i := range hashTable.array {
		hashTable.array[i] = &bucket{}
	}
	return hashTable

}

// Insert добавляет пару ключ-значение в хеш-таблицу
func (h *HashTable) Insert(key string, value int) {
	index := hash(key, len(h.array))
	h.array[index].insert(key, value)
}

// Search возвращает значение по ключу в хеш-таблице
func (h *HashTable) Search(key string) (int, bool) {
	index := hash(key, len(h.array))
	return h.array[index].search(key)
}

// Delete удаляет элемент по ключу из хеш-таблицы
func (h *HashTable) Delete(key string) {
	index := hash(key, len(h.array))
	h.array[index].delete(key)
}

// hash простая хеш-функция
func hash(key string, size int) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % size
}

// insert добавляет элемент в список
func (b *bucket) insert(key string, value int) {
	_, ok := b.search(key)
	if !ok {
		newNode := &bucketNode{key: key, value: value}
		newNode.next = b.head
		b.head = newNode
	} else {
		// Обновляем значение, если ключ уже существует
		currNode := b.head
		for currNode.key != key {
			currNode = currNode.next
		}
		currNode.value = value
	}
}

// search ищет ключ в списке
func (b *bucket) search(key string) (int, bool) {
	currNode := b.head
	for currNode != nil {
		if currNode.key == key {
			return currNode.value, true
		}
		currNode = currNode.next
	}
	return 0, false
}

// delete удаляет ключ из списка
func (b *bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}
	prevNode := b.head
	for prevNode.next.key != key {
		prevNode = prevNode.next
	}
	prevNode.next = prevNode.next.next
}
func (h *HashTable) Resize(newSize int) *HashTable {
	newHashTable := NewHashTable(newSize)
	for i := range h.array {
		currNode := h.array[i].head
		for currNode != nil {
			newHashTable.Insert(currNode.key, currNode.value)
			currNode = currNode.next
		}
	}
	return newHashTable
}
func (h *HashTable) Size() int {
	return len(h.array)
}

func main() {
	hashTable := NewHashTable(10)

	fmt.Println("Hash Table Created")
	fmt.Println("Hash Table Size:", hashTable.Size())
	hashTable.Insert("name", 123)
	hashTable.Insert("age", 456)
	value, _ := hashTable.Search("name")
	fmt.Println("name:", value)
	value, _ = hashTable.Search("age")
	fmt.Println("age:", value)

	hashTable.Delete("name")
	value, _ = hashTable.Search("name")
	fmt.Println("name:", value)

	hashTable = hashTable.Resize(50)
	fmt.Println("Hash Table Resize")
	fmt.Println("Hash Table Size:", hashTable.Size())

	value, _ = hashTable.Search("name")
	fmt.Println("name:", value)
	value, _ = hashTable.Search("age")
	fmt.Println("age:", value)

}
