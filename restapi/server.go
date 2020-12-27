package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// example
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

type Response struct {
	Code int 				`json: "code"`
	Message String 	`json: "message"`
	Data inteface{} 	`json: "data"`
}

type Resource interface {
	Uri() String
	Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) Response
	Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) Response
	Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) Response
	Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) Response
}

type (
	GetNotSupported struct{}
	PostNotSupported struct{}
	PutNotSupported struct{}
	DeleteNotSupported struct{}
)

func (GetNotSupported) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) Response {
	return Response{405, "", nil}
}

func (PostNotSupported) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) Response {
	return Response{405, "", nil}
}

func (PutNotSupported) Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) Response {
	return Response{405, "", nil}
}

func (DeleteNotSupported) Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) Response {
	return Response{405, "", nil}
}


func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
