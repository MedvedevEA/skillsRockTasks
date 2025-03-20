// Проведено измерение и сравнение двух алгоритмов сортировки
package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func bubbleSort(arr []int) ([]int, int) {
	n := len(arr)
	count := 0
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				count += 2
			}
		}
	}
	return arr, count
}
func insertionSort(arr []int) ([]int, int) {
	count := 0
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		count += 2

		// Перемещаем элементы arr[0..i-1], которые больше чем key, на одну позицию вперед
		// от текущей позиции
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
			count += 2
		}
		arr[j+1] = key
		count += 1
	}
	return arr, count
}

func main() {
	b123 := make([]int, 0, 10000)
	b321 := make([]int, 0, 10000)
	b231 := make([]int, 0, 10000)
	i123 := make([]int, 0, 10000)
	i321 := make([]int, 0, 10000)
	i231 := make([]int, 0, 10000)

	for i := 0; i < 10000; i++ {
		b123 = append(b123, i)
		b321 = append(b321, 9999-i)
		b231 = append(b231, rand.IntN(10000))
		i123 = append(i123, i)
		i321 = append(i321, 9999-i)
		i231 = append(i231, rand.IntN(10000))

	}

	var (
		start time.Time
		count int
	)
	start = time.Now()
	_, count = bubbleSort(b123)
	fmt.Printf("BUBBLESORT 123. COUNT: %d TIME: %v\n", count, time.Since(start))
	start = time.Now()
	_, count = bubbleSort(b321)
	fmt.Printf("BUBBLESORT 321. COUNT: %d TIME: %v\n", count, time.Since(start))
	start = time.Now()
	_, count = bubbleSort(b231)
	fmt.Printf("BUBBLESORT 231. COUNT: %d TIME: %v\n", count, time.Since(start))
	start = time.Now()
	_, count = insertionSort(i123)
	fmt.Printf("INSERTINONSORT 123. COUNT: %d TIME: %v\n", count, time.Since(start))
	start = time.Now()
	_, count = insertionSort(i321)
	fmt.Printf("INSERTINONSORT 321. COUNT: %d TIME: %v\n", count, time.Since(start))
	start = time.Now()
	_, count = insertionSort(i231)
	fmt.Printf("INSERTINONSORT 231. COUNT: %d TIME: %v\n", count, time.Since(start))

}
