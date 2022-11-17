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

	fs := http.FileServer(http.Dir("asset/"))
    http.Handle("/asset/", http.StripPrefix("/asset/", fs))

	http.ListenAndServe(":10000", nil)
}
