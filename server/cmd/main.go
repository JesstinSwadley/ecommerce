package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Name string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /products", handleProducts)

	fmt.Println("Server listening to Port 8080")
	http.ListenAndServe(":8080", mux)
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if product.Name == "" {
		http.Error(w, "product name is required", http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusCreated)
}
