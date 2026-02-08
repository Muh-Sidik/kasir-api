package route

import (
	"database/sql"
	"net/http"

	"github.com/Muh-Sidik/kasir-api/config"
	"github.com/Muh-Sidik/kasir-api/internal/handler"
	"github.com/Muh-Sidik/kasir-api/internal/repository"
	"github.com/Muh-Sidik/kasir-api/internal/service"
)

func ProductRoute(mux *http.ServeMux, e *config.Env, db *sql.DB) {
	handler := handler.NewProductHandler(
		service.NewProductService(
			repository.NewProductRepository(db),
		),
	)

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
}
