package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	get()
	client()
}

func get() {
	resp, err := http.Get("https://www.google.com")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", string(data))
}

func client() {
	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}

	// req.Header.Add("User-Agent", "LTS")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)
	fmt.Println(str)
}
