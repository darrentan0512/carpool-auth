package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is an initial endpoint")
}
func main() {

	http.HandleFunc("/api", Home)
	http.ListenAndServe(":8080", nil)
}
