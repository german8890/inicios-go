package main

import "fmt"

func main() {

	v := sumar(3, 4)
	fmt.Println(v)

	r := restar(5, 10)
	fmt.Println(r)
}

func sumar(v int, c int) int {
	return v + c
}

func restar(a int, b int) int {
	return a - b
}
