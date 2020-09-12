package main

import "fmt"

func main() {
	for 
}

func Piramid() {
	for i := 1; i < 6; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Gugudan() {
	for i := 2; i < 10; i++ {
		fmt.Printf("================\n%d 단\n", i)
		for j := 1; j < 10; j++ {
			if i == 3 && j == 2 {
				continue
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}
}
