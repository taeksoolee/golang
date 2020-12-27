package main

import (
	"net/http"
)

/* http.Handler interface
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/

type testHandler struct {
	http.Handler
}

func main() {
	http.Handle("/", new(testHandler))

	http.ListenAndServe(":5000", nil)
}

func (h * testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "Your Request Path is " + req.URL.Path
	w.Write([]byte(str))
}