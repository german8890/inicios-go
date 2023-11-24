package main

import "fmt"

func main() {

	v := sumar(3, 4)
	fmt.Println(v)
}

func sumar(v int, c int) int {
	return v + c
}
