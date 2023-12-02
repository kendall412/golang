package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	/* function handler will run a function when at the route.
	HandleFunc is a convenience method which takes the function and creates an http handler from it and adds it to default serveMux.
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "OOPS", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Data received: %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":9090", nil)
}
