package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Member struct {
	Name   string
	Age    int
	Active bool
}

type Members struct {
	Member []Member
}

func main() {
	// xmlEncoding()
	// xmlDecoding()
	xmlDecodingInFile()
}

func xmlEncoding() {
	mem := Member{"Alex", 10, true}

	xmlBytes, err := xml.Marshal(mem)
	if err != nil {
		panic(err)
	}

	xmlString := string(xmlBytes)
	fmt.Println(xmlString)
}

func xmlDecoding() {
	xmlBytes, _ := xml.Marshal(Member{"Time", 1, true})

	var mem Member
	err := xml.Unmarshal(xmlBytes, &mem)
	if err != nil {
		panic(err)
	}

	fmt.Println(mem.Name, mem.Age, mem.Active)
}

func xmlDecodingInFile() {
	fp, err := os.Open("c:\\temp\\test.xml")
	if err != nil {
		panic(err)
	}

	defer fp.Close()

	data, err := ioutil.ReadAll(fp)
	var members Members
	xmlerr := xml.Unmarshal(data, &members)
	if xmlerr != nil {
		panic(xmlerr)
	}

	fmt.Println(members)
}
