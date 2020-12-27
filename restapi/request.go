package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const API_KEY string = ""

var url string = "https:/www.naver.com"

func main() {

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
	dec := json.NewDecoder(os.Stdin)
}
