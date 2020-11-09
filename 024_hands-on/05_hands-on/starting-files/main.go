package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template
var err error

func init() {
	tmpl, err = template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics", fs)
	http.HandleFunc("/dogs", dogs)
	http.ListenAndServe(":8080", nil)
}

func dogs(res http.ResponseWriter, req *http.Request) {
	err := tmpl.Execute(res, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
