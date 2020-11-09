package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", home)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":8080", nil)
}

func home(writer http.ResponseWriter, reader *http.Request) {
	err := tmpl.Execute(writer, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
