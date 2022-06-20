package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful")
	//getting values from the filled form
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name  = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//r  is request. w is Responnse
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")

}

func main() {
	//telling Golang to chechk the static directory
	fileServer := http.FileServer(http.Dir("./static"))
	//to handle, send to fileServer to serve index file
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	//creating server  with listeand serve. One of this two values will be assigned toe rr
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
