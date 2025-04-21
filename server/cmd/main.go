package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

	mux.HandleFunc("POST /products", handleProducts)

	ConnectDatabase()

	fmt.Printf("Server listening to Port: %v", port)
	http.ListenAndServe(port, mux)
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
