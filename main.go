package main

import (
	"fmt"
	"os"
)

func main() {
	var a int
	var b int
	fmt.Println("Введите значение переменной а ")
	fmt.Fscan(os.Stdin, &a)

	fmt.Println("Введите значение переменной b ")
	fmt.Fscan(os.Stdin, &b)

	if a > b {
		fmt.Println("a больше b")
	} else if a < b {
		fmt.Println("a меньше b")
	} else {
		fmt.Println("a равняется b")
	}
}
