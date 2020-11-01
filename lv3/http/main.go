package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	Exec3()
}

func Exec1() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello world"))
	})

	http.ListenAndServe(":5000", nil)
}

type testHandler struct {
	http.Handler
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "Your Request Path is " + req.URL.Path
	w.Write([]byte(str))
}

func Exec2() {
	http.Handle("/", new(testHandler))

	http.ListenAndServe(":5000", nil)
}

type staticHandler struct {
	http.Handler
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	localPath := "wwwroot" + req.URL.Path
	fmt.Print(localPath)
	content, err := ioutil.ReadFile(localPath)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}

	w.Write(content)
}

func Exec3() {
	http.Handle("/", new(staticHandler))

	http.ListenAndServe(":5000", nil)
}
