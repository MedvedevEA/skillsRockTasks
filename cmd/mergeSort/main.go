package main

import (
	"fmt"
)

// Функция для слияния двух подмассивов
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Слияние двух отсортированных подмассивов в один
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Добавление оставшихся элементов, если они есть
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// Рекурсивная функция сортировки слиянием
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func main() {
	sample := []int{9, -3, 5, 2, 6, 8, -6, 1, 3}
	sorted := mergeSort(sample)
	fmt.Println("Отсортированный массив:", sorted)
	sample2 := []int{}
	sorted2 := mergeSort(sample2)
	fmt.Println("Отсортированный массив:", sorted2)
	sample3 := []int{0}
	sorted3 := mergeSort(sample3)
	fmt.Println("Отсортированный массив:", sorted3)
	sample4 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0}
	sorted4 := mergeSort(sample4)
	fmt.Println("Отсортированный массив:", sorted4)

}
