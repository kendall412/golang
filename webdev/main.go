package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {

	// tpl, _ = template.ParseGlob("templates/*.html")
	tpl, _ = template.ParseFiles("templates/index")

	http.HandleFunc("/", indexHandler)

	/* http.ListenAndServe:

	implicitly knows :8080 really means localhost:8080

	entering "nil" implicitly uses DefaultServeMux
	ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls teh handler for the pattern that most closely matches the URL.*/

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// f.Fprint(w, "Hello, Danny!")
	tpl.Execute(w, nil)
}
