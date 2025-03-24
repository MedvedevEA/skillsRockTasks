// Изменена функция quickSort. Выбор пивотом последнего элемента массива позволило упростить код.
package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func main() {
	sample := []int{9, -3, 5, 2, 6, 8, -6, 1, 3}
	sorted := quickSort(sample)
	fmt.Println("Отсортированный массив:", sorted)
	sample2 := []int{}
	sorted2 := quickSort(sample2)
	fmt.Println("Отсортированный массив:", sorted2)
	sample3 := []int{0}
	sorted3 := quickSort(sample3)
	fmt.Println("Отсортированный массив:", sorted3)

}
