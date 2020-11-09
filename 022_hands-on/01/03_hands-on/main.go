package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(res, "template.gohtml", "Awesome")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
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
