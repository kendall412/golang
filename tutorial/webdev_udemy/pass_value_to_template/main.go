package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

var pfile string
pfile = "template/one.gohtml"

var data string
data = "obeying God"

func init(){
	tpl = template.Must(template.ParseFiles(pfile))
}

func main() {
err:= tpl.ExecuteTemplate(os.Stdout, pfile, data)
if err != nil{
	log.Fatalln(err)
}

}
