package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		fmt.Println("There is no server port defined")
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /products", handleCreateProducts)

	err := http.ListenAndServe(":"+port, mux)

	fmt.Println(err)

	if err != nil {
		log.Fatal(err)
	}
}

func handleCreateProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
