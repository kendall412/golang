/*
see text/template in Go standard doc
https://pkg.go.dev/text/template
*/

package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	html_file := "tpl.gohtml"
	/*
		func ParseFiles(filenames ...string)(*Template, error)
	*/
	tpl, err := template.ParseFiles(html_file)
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	/*
		returns *os.File
		func (f *File) Write(b []byte) (n int, err error)
	*/
	if err != nil {
		log.Println(err)
	}

	defer nf.Close()

	/*
		func (t *Template) Execute(wr io.Writer, data any) error

		type Writer interface {
			Write(p []byte) (n int, err error)
		}
	*/
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err2 := tpl.Execute(nf, nil)
	if err2 != nil {
		log.Fatalln(err2)
	}

}
