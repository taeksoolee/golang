package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) PrintName() {
	fmt.Println(p.name)
}

func PersonTest() {
	// var p person
	p1 := Person{"개똥이", 15}
	p2 := Person{name: "소똥이", age: 21}
	p3 := Person{name: "Carson"}
	p4 := Person{}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)

	p1.name = "깨똥이"
	fmt.Println(p1)

	p1.PrintName()
}

type Student struct {
	name    string
	class   int
	sungjuk Sungjuk
}

type Sungjuk struct {
	name  string
	grade string
}

func (s Student) ViewSunguk() {
	fmt.Println(s.sungjuk)
}

func (s *Student) InputSunguk(name, grade string) {
	s.sungjuk.name = name
	s.sungjuk.grade = grade
}

func StudentTest() {
	var s Student
	s = Student{name: "leetaeksoo", class: 1}

	s.InputSunguk("수학", "B")
	s.ViewSunguk()
}

func main() {
	// PersonTest()
	StudentTest()
}
