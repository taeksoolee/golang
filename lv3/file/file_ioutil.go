package main

import "io/ioutil"

func main() {
	test2()

}

func test1() {
	bytes, err := ioutil.ReadFile("C:\\temp\\1.txt")

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("C:\\temp\\2.txt", bytes, 0)

	if err != nil {
		panic(err)
	}
}

func test2() {
	bytes, err := ioutil.ReadFile("src/wwwroot/index.html")

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("src/wwwroot/index2.html", bytes, 0)

	if err != nil {
		panic(err)
	}
}