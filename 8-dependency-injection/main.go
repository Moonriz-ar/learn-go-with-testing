package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	// fmt.Fprintf is like fmt.Printf but instead takes a writer to send the strings to, whereas fmt.Printf defaults to stdout
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5501", http.HandlerFunc(MyGreeterHandler)))
}
