package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// post()
	postForm()
}

// Post
func post() {
	reqBody := bytes.NewBufferString("Post Plain text")
	resp, err := http.Post("http://httpbin.org/post", "text/plain", reqBody)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		print(str)
	}
}

// Post form
func postForm() {
	resp, err := http.PostForm("http://httpbin.org/post", url.Values{"Name": {"Lee"}, "Age": {"10"}})
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

// Post JSON
func postJson() {
	person := Person{"Alex", 10}
	pbytes, _ := json.Marshal(person)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://httpbin.org/post", "application/json", buff)

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

// Post XML
func postXML() {
	person := Person{"Alex", 10}
	pbytes, _ := xml.Marshal(person)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://httpbin.org/post", "applcation/xml", buff)

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

// client
func client() {
	person := Person{"Alex", 10}
	pbyes, _ := xml.Marshal(person)
	buff := bytes.NewBuffer(pbytes)

	req, err := http.NewRequest("POST", "http://httpbin.org/post", buff)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/xml")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}
