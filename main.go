package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Muh-Sidik/kasir-api/database"
	"github.com/Muh-Sidik/kasir-api/docs"
	_ "github.com/Muh-Sidik/kasir-api/docs"
	"github.com/Muh-Sidik/kasir-api/internal/handler"
	"github.com/swaggo/http-swagger/v2"
)

// @title Swagger Kasir API
// @version 1.0
// @description This is a kasir server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host kasir-api-production-e286.up.railway.app
// @BasePath /
func main() {
	mode := "local"

	switch mode {
	case "local":
		docs.SwaggerInfo.Host = "localhost:8000"
	case "railway":
		docs.SwaggerInfo.Host = "kasir-api-production-e286.up.railway.app"
	}
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := database.New()
	defer db.Close()

	handler := handler.Handler{
		DB: db,
	}

	mux := http.NewServeMux()

	// DELETE http://localhost:8000/api/categories/{id}
	mux.HandleFunc("DELETE /api/categories/{id}", handler.DeleteCategoryByID)
	// PUT http://localhost:8000/api/categories/{id}
	mux.HandleFunc("PUT /api/categories/{id}", handler.UpdateCategoryByID)
	// GET http://localhost:8000/api/categories/{id}
	mux.HandleFunc("GET /api/categories/{id}", handler.GetCategoryByID)

	// POST http://localhost:8000/api/categories
	mux.HandleFunc("POST /api/categories", handler.CreateCategory)
	// GET http://localhost:8000/api/categories
	mux.HandleFunc("GET /api/categories", handler.Categories)

	// DELETE http://localhost:8000/api/product/{id}
	mux.HandleFunc("DELETE /api/product/{id}", handler.DeleteProductByID)
	// PUT http://localhost:8000/api/product/{id}
	mux.HandleFunc("PUT /api/product/{id}", handler.UpdateProductByID)
	// GET http://localhost:8000/api/product/{id}
	mux.HandleFunc("GET /api/product/{id}", handler.GetProductByID)

	// POST http://localhost:8000/api/product
	mux.HandleFunc("POST /api/product", handler.CreateProduct)
	// GET http://localhost:8000/api/product
	mux.HandleFunc("GET /api/product", handler.Products)

	mux.HandleFunc("GET /docs/", httpSwagger.Handler(
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))

	// GET http://localhost:8000/health
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Server is running",
		})
	})

	// GET http://localhost:8000
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Kasir Api, check documentation at " + r.Host + "/docs",
		})
	})

	fmt.Println("Successfully listen server in port :8000")
	err := http.ListenAndServe(
		":8000",
		mux,
	)

	if err != nil {
		log.Fatalf("error server: %v", err)
	}

}
