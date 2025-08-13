package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Product struct {
	Name  string
	Price float64
}

func main() {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		fmt.Println("There is no server port defined")
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /products", handleCreateProducts)
	mux.HandleFunc("GET /products", handleListAllProducts)

	err := http.ListenAndServe(":"+port, mux)

	if err != nil {
		log.Fatal(err)
	}
}

func handleCreateProducts(w http.ResponseWriter, r *http.Request) {
	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	fmt.Println(product)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Product was created"))
}

func handleListAllProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("List All Products"))
}
