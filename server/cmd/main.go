package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type Product struct {
	Name string
}

var Db *sql.DB

func main() {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		fmt.Println("There is no server port defined")
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /products", handleCreateProducts)
	mux.HandleFunc("GET /products", handleGetAllProducts)

	ConnectDatabase()

	fmt.Printf("Server listening to Port: %v", port)
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

	if product.Name == "" {
		http.Error(w, "product name is required", http.StatusBadRequest)

		return
	}

	_, err = Db.Exec("INSERT INTO products (name) values ($1)", product.Name)

	if err != nil {
		fmt.Println(err)

		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}

func handleGetAllProducts(w http.ResponseWriter, r *http.Request) {
	var product string
	var products []string

	rows, err := Db.Query("SELECT * FROM prodcuts")

	defer rows.Close()

	if err != nil {
		fmt.Println(err)

		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	for rows.Next() {
		rows.Scan(&product)
		products = append(products, product)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(product))
}

func ConnectDatabase() {
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	psqlConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	db, errSql := sql.Open("postgres", psqlConn)

	if errSql != nil {
		fmt.Println("There is an error while connecting to the database", errSql)
	} else {
		Db = db
		fmt.Println("Succesfully connected to database!")
	}
}
