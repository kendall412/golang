package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template
var shinoo hurs

type hurs struct {
	name       string
	age        int
	occupation string
}

func main() {
	shinoo = hurs{
		name:       "Shinoo",
		age:        11,
		occupation: "kid"}

	tpl, _ = template.ParseGlob("templates/*")
	http.HandleFunc("/hurs", hursHandler)
	http.ListenAndServe(":8080", nil)
}

func hursHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "hurs", shinoo)
}
