package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Parse form error %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful \n")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "The name is %s", name)
	fmt.Fprintf(w, "The address is %s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not suited", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")

}

func main() {

	//telling golang to check the static directory
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver) //index.html file
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8082\n")

	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}

}
