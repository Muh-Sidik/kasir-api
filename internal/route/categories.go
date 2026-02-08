package route

import (
	"database/sql"
	"net/http"

	"github.com/Muh-Sidik/kasir-api/config"
	"github.com/Muh-Sidik/kasir-api/internal/handler"
	"github.com/Muh-Sidik/kasir-api/internal/repository"
	"github.com/Muh-Sidik/kasir-api/internal/service"
)

func CategoryRoute(mux *http.ServeMux, e *config.Env, db *sql.DB) {
	handler := handler.NewCategoryHandler(
		service.NewCategoryService(
			repository.NewCategoryRepository(db),
		),
	)

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
}
