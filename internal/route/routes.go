package route

import (
	"database/sql"
	"net/http"

	"github.com/Muh-Sidik/kasir-api/config"
	"github.com/Muh-Sidik/kasir-api/internal/handler"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/response"
	"github.com/Muh-Sidik/kasir-api/internal/repository"
	"github.com/Muh-Sidik/kasir-api/internal/service"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Setup(mux *http.ServeMux, e *config.Env, db *sql.DB) {
	handler := &handler.Handler{
		ProductSrv: service.NewProductService(
			repository.NewProductRepository(db),
		),
		CategorySrv: service.NewCategoryService(
			repository.NewCategoryRepository(db),
		),
	}

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
		response.OK(
			"Server is running",
			nil,
			nil,
		).JSON(w, http.StatusOK)
	})

	// GET http://localhost:8000
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		response.OK(
			"Kasir Api, check documentation at "+r.Host+"/docs",
			nil,
			nil,
		).JSON(w, http.StatusOK)
	})
}
