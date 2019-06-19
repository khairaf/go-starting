package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Whoaa,, look at this!")
}

func about_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Ini adalah an about page")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil) //nill = nothing to pass as server argument
}