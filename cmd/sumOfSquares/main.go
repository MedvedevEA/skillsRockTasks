// Реализован свой вариант функции для вычисления суммы квадратов. Добавлены асинхронные вычисления и функция отмены.
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func SumOfSquares(ctx context.Context, number int) (int, error) {
	if number < 0 {
		return 0, errors.New("number must be greater than zero")
	}
	sum := 0
	ch := make(chan int, number)
	for i := 1; i <= number; i++ {
		go func(ch chan int, number int) {
			ch <- (number * number)
		}(ch, i)
	}
	for i := 1; i <= number; i++ {
		select {
		case <-ctx.Done():
			return 0, errors.New("function canceled")
		case value := <-ch:
			sum += value
		}
	}
	close(ch)
	return sum, nil
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//Отмена контекста после таймаута
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()
	sum, err := SumOfSquares(ctx, 5)

	fmt.Println(sum, err)

}
