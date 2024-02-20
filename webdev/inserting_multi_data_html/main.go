package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template
var shinoo hurs

/*
Data must be exportable (meaning must be capitalized) in order to using in template. Alternative is to use accessor (https://stackoverflow.com/questions/21863760/accessing-struct-variable-in-slice-of-many-structs-in-html-template-golang).
*/

type hurs struct {
	Name       string
	Age        int
	Occupation string
}

func main() {
	shinoo = hurs{
		Name:       "Shinoo",
		Age:        11,
		Occupation: "kid"}

	tpl, _ = template.ParseGlob("templates/*")
	http.HandleFunc("/hurs", hursHandler)
	http.ListenAndServe(":8080", nil)
}

func hursHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "hurs", shinoo)
}
