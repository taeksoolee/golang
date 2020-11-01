package main

import "fmt"

// 함수 호출시 복사 됨.
func add(x int, y int) int {
	return x + y
}

func fun1(x, y int) (int, int) {
	return y, x
}

func fun2(x, y int) {
	fmt.Println(y, x)
}

func FuncTest() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d + %d = %d\n", 1, i, add(1, i))
	}

	a, b := fun1(2, 3)

	fmt.Println(a, b)

	fun2(2, 3)
}

func f1(x int) {
	if x == 0 {
		return
	}
	fmt.Printf("before %d \n", x)
	f1(x - 1)
	fmt.Printf("after %d \n", x)
}

func sum(x, s int) int {
	if x == 0 {
		return s
	}

	s += x
	return sum(x-1, s)
}

func main() {

}
