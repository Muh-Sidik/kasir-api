package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

var product = []*Produk{
	{1, "Indomie Godog", 3500, 10},
	{2, "Vit 1000ml", 3000, 40},
}

func getProductByID(id int, w http.ResponseWriter) {
	for _, p := range product {
		if p.ID == id {
			json.NewEncoder(w).Encode(*p)
			return
		}
	}

	http.Error(w, "Produk Not Found", http.StatusNotFound)
}

func deleteProductByID(id int, w http.ResponseWriter) {
	for i, p := range product {
		if p.ID == id {
			productName := p.Nama
			product = append(product[:i], product[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "OK",
				"message": "Success delete produk " + productName,
			})
			return
		}
	}

	http.Error(w, "Produk Not Found", http.StatusNotFound)
}

func updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	updateProduct := new(Produk)
	err := json.NewDecoder(r.Body).Decode(updateProduct)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	for i := range product {
		if product[i].ID == id {
			updateProduct.ID = product[i].ID
			product[i].Nama = updateProduct.Nama
			product[i].Harga = updateProduct.Harga
			product[i].Stok = updateProduct.Stok

			json.NewEncoder(w).Encode(*updateProduct)
			return
		}
	}
	http.Error(w, "Produk Not Found", http.StatusNotFound)
}

func main() {
	// DELETE http://localhost:8000/api/produk/:id
	// PUT http://localhost:8000/api/produk/:id
	// GET http://localhost:8000/api/produk/:id
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		idParam := strings.TrimPrefix(r.URL.Path, "/api/produk/")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			http.Error(w, "Invalid Produk Id", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case "GET":
			getProductByID(id, w)
		case "PUT":
			updateProduct(id, w, r)
		case "DELETE":
			deleteProductByID(id, w)
		}
	})

	// POST http://localhost:8000/api/produk
	// GET http://localhost:8000/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "POST" {
			newProduct := new(Produk)
			err := json.NewDecoder(r.Body).Decode(newProduct)

			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}

			newProduct.ID = len(product) + 1
			product = append(product, newProduct)

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(*newProduct)
			return
		}

		json.NewEncoder(w).Encode(product)
	})

	// GET http://localhost:8000/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Kasih is running",
		})
	})

	fmt.Println("Successfully listen server in port :8000")
	err := http.ListenAndServe(
		":8000",
		nil,
	)

	if err != nil {
		log.Fatalf("error server: %v", err)
	}

}
