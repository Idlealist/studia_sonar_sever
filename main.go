package main

import (
    "net/http"
    "encoding/json"
    "log"
)

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

var products = []Product{
    {ID: 1, Name: "Laptop", Price: 4999.99},
    {ID: 2, Name: "TV", Price: 2999.99},
    {ID: 3, Name: "Smartwatch", Price: 699.99},
}

func main() {
    http.HandleFunc("/products", handleProducts)
    http.HandleFunc("/checkout", handleCheckout)

    if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start %v", err)
	}
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
    setupCORS(&w, r)
    if r.Method == http.MethodOptions {
        return
    }
    if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Failed to encode products", http.StatusInternalServerError)
		log.Printf("Encode error %v", err)
	}
}

func handleCheckout(w http.ResponseWriter, r *http.Request) {
    setupCORS(&w, r)
    if r.Method == http.MethodOptions {
        return
    }
    var order []Product
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid order format", http.StatusBadRequest)
		log.Printf("Decode error %v", err)
		return
	}
    log.Println("Order:", order)
    w.WriteHeader(http.StatusOK)
}

func setupCORS(w *http.ResponseWriter, r *http.Request) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    (*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
