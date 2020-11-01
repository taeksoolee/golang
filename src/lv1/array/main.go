package main

import "fmt"

func main() {
	// var A [10]int
	copyArray()
	// reverseArray()
	reverseArray2()
	sortArray()
}

func sortArray() {
	// redix sort

	arr := [11]int{0, 5, 4, 9, 1, 2, 8, 3, 6, 4, 5}
	temp := [11]int{}

	for i := 0; i < len(arr); i++ {
		idx := arr[i]
		temp[idx]++
	}

	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr[idx] = i
			idx++
		}
	}

	fmt.Println(arr)
}

func reverseArray2() {
	// N/2
	arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	fmt.Println(arr)
}

func reverseArray() {
	// 2N
	arr := [5]int{1, 2, 3, 4, 5}
	temp := [5]int{}

	for i := 0; i < 5; i++ {
		temp[i] = arr[len(arr)-i-1]
	}

	for i := 0; i < 5; i++ {
		arr[i] = temp[i]
	}

	fmt.Println(temp)
	fmt.Println(arr)
}

func copyArray() {
	arr := [5]int{1, 2, 3, 4, 5}
	clone := [5]int{}

	for i := 0; i < len(arr); i++ {
		clone[i] = arr[i]
	}

	fmt.Println(clone)
}

func goString() {
	// string >> rune[], byte[]
	// UTF8 >> 1 ~ 3 byte

	// s := "hello 월드"

	// s2 := []rune(s)

	var s3 []rune
	s3 = []rune("안녕 세상")

	for i := 0; i < len(s3); i++ {
		fmt.Print(string(s3[i]), ", ")
	}
}
