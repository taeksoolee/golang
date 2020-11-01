package main

import "fmt"

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("최종 i 값 : ", i)

	for {
		i++
		fmt.Println(i)
	}
}

func Switch() {
	x := 33

	switch x {
	case 31:
		fmt.Print("x = 31\n")
	case 32:
		fmt.Print("x = 32\n")
	case 33:
		fmt.Print("x = 33\n")
	}

	switch {
	case x > 40:
		fmt.Println("x는 40보다 크다.")
	case x < 20:
		fmt.Println("x는 20보다 작다.")
	case x > 30:
		fmt.Println("x는 30보다 크다")
	}
}
