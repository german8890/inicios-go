package main

import "fmt"

func main() {

	v := sumar(3, 4)
	fmt.Println(v)

	r := restarDiego(5, 10)
	fmt.Println(r)
	m:=Miltiplicate (8, 6)
	fmt.Println(m)
}

func sumar(v int, c int) int {
	return v + c
}

func restarDiego(a int, b int) int {
	return a - b
}
func Miltiplicate(a int, b int) int {
	return a * b
}