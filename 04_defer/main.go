package main

import "fmt"

// F ...
func F() {
	panic("panic")
}

func main() {

	// 1. Порядок вывода defer

	defer fmt.Println("Вызов первого defer")
	defer fmt.Println("Вызов второго defer")

	// 2. Вычисление аргументов defer

	value := 5

	defer fmt.Println("Инкремент value со значением 5 = ", value)

	value++

	defer func() {
		fmt.Println("Инкремент value со значением 5 через замыкание = ", value)
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Паника перехвачена:", r)
		}
	}()

	panic("ошибка")
}
