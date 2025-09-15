package main

import (
	"fmt"
)

func find[T comparable](array []T, needle T) (int, bool) {
	for i, v := range array {
		if v == needle {
			return i, true
		}
	}

	return -1, false
}

func reverse(a []int) []int {
	for f, l := 0, len(a)-1; f < len(a)/2; f, l = f+1, l-1 {
		a[f], a[l] = a[l], a[f]
	}

	return a
}

func main() {
	// 1. Создать массив из 5 целых чисел и заполнить его значениями от 1 до 5.

	a := [5]int{1, 2, 3, 4, 5}

	// 2. Вывести массив в цикле `for`.

	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// 3. Сложить все элементы массива и вывести сумму.

	sum := 0
	for _, v := range a {
		sum += v
	}

	fmt.Println("Sum of array:", sum)

	// 4. Найти минимальное и максимальное значение в массиве.

	minValue := a[0]
	maxValue := a[0]

	for _, v := range a[1:] {
		minValue = min(minValue, v)
		maxValue = max(maxValue, v)
	}

	fmt.Printf("Min of array: %d, max of array: %d\n", minValue, maxValue)

	// 5. Считать `n` чисел от пользователя и сохранить их в массив (если n фиксированное).
	n := 5

	a1 := [5]int{}

	for i := range n {
		var val int
		fmt.Scanln(&val)
		a1[i] = val
	}

	for _, v := range a1 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// 6. Реализовать функцию поиска элемента в массиве (линейный поиск).

	a2 := []int{1, 2, 3, 4, 5}

	if f, ok := find(a2, 3); ok != false {
		fmt.Println("Find index of 3 in array (1, 2, 3, 4, 5)", f)
	}

	// 7. Развернуть массив

	a3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Разворачиваем массив ")
	for _, v := range reverse(a3) {
		fmt.Printf("%d ", v)
	}

	fmt.Printf("\nРезультат ")

	for _, v := range reverse(a3) {
		fmt.Printf("%d ", v)
	}
}
