package main

import (
	"fmt"
)

// binarySearch функция для выполнения бинарного поиска в отсортированном массиве.
// Возвращает индекс целевого значения, если оно присутствует в массиве, иначе возвращает -1.
func binarySearch[T int | string](arr []T, target T) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2 // Используем (high-low)/2 для избежания переполнения

		// Проверяем, равен ли средний элемент целевому значению
		if arr[mid] == target {
			return mid
		}

		// Если целевое значение больше, игнорируем левую половину
		if arr[mid] < target {
			low = mid + 1
		} else {
			// Если целевое значение меньше, игнорируем правую половину
			high = mid - 1
		}
	}

	// Если элемент не найден, возвращаем -1
	return -1
}

func main() {
	//1
	arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target := 23
	result := binarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Элемент найден на позиции: %d\n", result)
	} else {
		fmt.Println("Элемент не найден в массиве.")
	}
	//2
	arr2 := []string{"a", "b", "c", "d", "e", "f", "j", "h", "i", "j"}
	target2 := "b"
	result2 := binarySearch(arr2, target2)
	if result2 != -1 {
		fmt.Printf("Элемент найден на позиции: %d\n", result2)
	} else {
		fmt.Println("Элемент не найден в массиве.")
	}
	//3
	arr3 := []string{}
	target3 := "a"
	result3 := binarySearch(arr3, target3)
	if result3 != -1 {
		fmt.Printf("Элемент найден на позиции: %d\n", result3)
	} else {
		fmt.Println("Элемент не найден в массиве.")
	}
	//4
	arr4 := []string{"b", "c"}
	target4 := "a"
	result4 := binarySearch(arr4, target4)
	if result4 != -1 {
		fmt.Printf("Элемент найден на позиции: %d\n", result4)
	} else {
		fmt.Println("Элемент не найден в массиве.")
	}

}
