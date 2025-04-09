package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /products", handleProducts)

	fmt.Println("Server listening to Port 8080")
	http.ListenAndServe(":8080", mux)
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from products")
}
