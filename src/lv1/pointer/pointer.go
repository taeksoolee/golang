package main

import "fmt"

func GetPointer() (*int, int) {
	var n int
	n = 1
	return &n, n
}

func PointerTest() {
	p1, v1 := GetPointer()
	p2, v2 := GetPointer()

	fmt.Printf("1 >> %d\t%p\n", v1, p1)
	fmt.Printf("2 >> %d\t%p\n", v2, p2)
}

type Student struct {
	name string
	age  int

	grade string
}

func (s Student) PrintStudent() {
	// value >> s의 전체가 복사
	fmt.Printf("%s\t%d\t%s", s.name, s.age, s.grade)
}

func (s *Student) InputStudent(grade string) {
	// pointer >> s의 주소만 복사됨
	s.grade = grade
}

func main() {
	PointerTest()
	// c또는 c++에서는 	p -> attr 	/ v.attr
	// go에서는 		p.attr 		/ v.attr
}
