package main

import (
	"fmt"
	"log"
	"net/http"
)

// localhost:8080/form.html With Choosen Name and Address Then Click Submit
// Parse localhost:8080/form.html With Written Name and Address Then Redirecting to localhost:8080/form
// err = Error From Parse Result or Null
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Error Handling
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v\n", err)
	}
	if r.URL.Path != "/form" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "POST Request Successful!!!!\n")

	// form.html Variables
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Error Handling
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is Not Supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	// FileServer for / and /form and /hello
	// fileServer for / is index.html (Default)
	fileServer := http.FileServer(http.Dir("./html/"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting Local Static Simple Web Server at PORT 8080\n")
	// Create Server - err = Error or Null
	err := http.ListenAndServe(":8080", nil)
	// Error Handling
	if err != nil {
		log.Fatal(err)
	}
}
