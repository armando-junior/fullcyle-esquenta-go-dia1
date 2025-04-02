package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func getDBConnection() (*sql.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	log.Printf("Attempting to connect to database with connection string: %s", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error opening database connection: %v", err)
		return nil, err
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Try to ping the database
	for i := 0; i < 5; i++ {
		err = db.Ping()
		if err == nil {
			log.Printf("Successfully connected to database")
			return db, nil
		}
		log.Printf("Attempt %d: Error pinging database: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to database after 5 attempts")
}

func listProducts(w http.ResponseWriter, r *http.Request) {
	db, err := getDBConnection()
	if err != nil {
		log.Printf("Database connection error: %v", err)
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, description, price FROM products")
	if err != nil {
		log.Printf("Error querying products: %v", err)
		http.Error(w, "Error querying products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
		if err != nil {
			log.Printf("Error scanning product row: %v", err)
			http.Error(w, "Error scanning products", http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World, Armando!")
	})

	http.HandleFunc("/products", listProducts)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
