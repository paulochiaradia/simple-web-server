package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//check if the current request URL Path matches "/"
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}

	//Using r.Method to check whether the request is using GET or not
	if r.Method != "GET" {
		r.Header.Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("hello!"))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	//check the parsed form
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	//print the form data
	fmt.Fprintf(w, "POST request successful!")
	name := r.FormValue("name")
	address := r.FormValue("address")
	w.Write([]byte("\n"))
	w.Write([]byte(name))
	w.Write([]byte(address))

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
