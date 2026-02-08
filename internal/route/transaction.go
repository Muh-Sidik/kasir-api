package route

import (
	"database/sql"
	"net/http"

	"github.com/Muh-Sidik/kasir-api/config"
	"github.com/Muh-Sidik/kasir-api/internal/handler"
	"github.com/Muh-Sidik/kasir-api/internal/repository"
	"github.com/Muh-Sidik/kasir-api/internal/service"
)

func TransactionRoute(mux *http.ServeMux, e *config.Env, db *sql.DB) {
	handler := handler.NewTransactionHandler(
		service.NewTransactionService(
			repository.NewTransactionRepository(db),
		),
	)

	// POST http://localhost:8000/api/checkout
	mux.HandleFunc("POST /api/checkout", handler.HandleCheckout)
}
