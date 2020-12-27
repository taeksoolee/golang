package main

import (
	"encoding/json"
	"fmt"
)

type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	// jsonEncoding()
	jsonDecoding()
}

func jsonEncoding() {
	mem := Member{"Alex", 10, true}

	jsonBytes, err := json.Marshal(mem)

	if err != nil {
		panic(err)
	}

	jsonString := string(jsonBytes)
	fmt.Println(jsonString)
}

func jsonDecoding() {
	jsonBytes, _ := json.Marshal(Member{"Tim", 1, true})

	var mem Member
	err := json.Unmarshal(jsonBytes, &mem)
	if err != nil {
		panic(err)
	}

	fmt.Println(mem.Name, mem.Age, mem.Active)
}
