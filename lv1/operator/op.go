package main

import "fmt"

func main() {
	a := 2
	b := 4

	fmt.Printf("a&b = %v\n", (a & b))
	fmt.Printf("a|b = %v\n", (a | b))
	fmt.Println("a^b = ", a^b)
	fmt.Println(^a)
	fmt.Println("clear >> ", a&^a)
	fmt.Println("clear >> ", b&^b)

	c := 21
	d := c % 10
	e := c / 10
	fmt.Printf("%v, %v\n", d, e)

	f := 56             // 00001100
	fmt.Println(f << 1) // *2
	fmt.Println(f >> 1) // /2

	g := 3
	g += 1
	fmt.Println(g)
	g++
	fmt.Println(g)

	var h bool
	h = true
	fmt.Println(!h)
}
