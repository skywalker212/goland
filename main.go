package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I did not hit her, I DID NOT!")
}

func health_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "80 health, 20 humor, 5 dread")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/health", health_handler)
	http.ListenAndServe(":8000", nil)
}
