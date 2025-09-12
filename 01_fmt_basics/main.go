package main

import (
	"bufio"
	"fmt"
	"os"
)

// User — пример структуры для форматирования
type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Hello, Go!")

	// Пример строки и числа
	substr := "substing"
	num := 10

	// %s — строка, %d — число
	fmt.Printf("This is %s with number %d\n", substr, num)

	// %T — тип, %b — двоичное, %x/%X — шестнадцатеричное
	fmt.Printf("type (%T), Integer (%d), binary (%b), hexa (%x)(%X)\n", num, num, num, num, num)

	// Работа с float
	pi := 3.14159265
	fmt.Printf("Float (%.23f)\n", pi) // ограничение знаков после запятой

	// Считывание одного числа
	var n int
	fmt.Scanln(&n)
	fmt.Println("Product number * 2:", n*2)

	// Чтение строки с пробелами через bufio.Reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Printf("You typed: %s", text)

	// Пример структуры и разные форматы вывода
	user := User{"Satoshi", 42}
	// %v — обычный вид, %+v — с названиями полей, %#v — Go-вид
	fmt.Printf("user: %v\n%+v\n%#v\n", user, user, user)
}
