package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Product struct {
	Name string
}

func main() {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		fmt.Println("There is no server port defined")
		os.Exit(1)
	}

	ConnectDatabase()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /products", handleProducts)

	fmt.Printf("Server listening to Port %v", port)
	http.ListenAndServe(port, mux)
}

func ConnectDatabase() {
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Need to error handle all envs at some point
	// reference how it is done in main

	psqlConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := pgxpool.New(context.Background(), psqlConn)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully connected to database!")
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

	_, err = Db.Exec("insert into products (name) values ($1)", product.Name)

	if err != nil {
		fmt.Println(err)

		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}
