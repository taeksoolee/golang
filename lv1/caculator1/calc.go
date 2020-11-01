package main

import "fmt"

// int >> int(32bit 4 or 64bit 8 byte), int8, int16, int32, int64, uint16
// float >> float32, float64
// string >> rune[] (rune : 1~3byte)

// variable
// - attribute : name, value, type (size), memory address (start point)
func main() {
	// declare
	var a int
	var b int

	// assignment
	a = 3
	b = 4

	fmt.Println(a + b)

}
