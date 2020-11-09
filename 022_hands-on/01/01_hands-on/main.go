package main

import (
	"io"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Home")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dog")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Noyan")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
