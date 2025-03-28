// Реализован свой вариант функции по генерации пермутаций символов строки
package main

import (
	"fmt"
)

func f(n int) int {
	if n < 2 {
		return 1
	}
	return n * f(n-1)
}
func permute(str string) []string {
	if len(str) == 0 {
		return []string{}
	}
	if len(str) == 1 {
		return []string{str}
	}

	result := make([]string, 0, f(len(str)))

	for i := range str {
		list := permute(str[:i] + str[i+1:])
		for j := range list {
			result = append(result, string(str[i])+list[j])
		}
	}

	return result
}

func main() {
	fmt.Println(f(3))
	str := "1234"
	p := permute(str)
	fmt.Println("Permutations of", str, ":")
	for i := range p {
		fmt.Println(p[i])
	}
}
