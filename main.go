package main

import (
	"fmt"
	"net/http"
)

func topPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!!")
	fmt.Println("Endpoint Hit")
}

func main() {
	http.HandleFunc("/", topPage)
	http.ListenAndServe(":10000", nil)
}
