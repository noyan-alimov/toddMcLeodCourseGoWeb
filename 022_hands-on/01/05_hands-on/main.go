package main

import (
	"io"
	"net/http"
)

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dog")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Noyan")
}

func main() {
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}
