package route

import (
	"database/sql"
	"net/http"

	"github.com/Muh-Sidik/kasir-api/config"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/response"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Setup(mux *http.ServeMux, e *config.Env, db *sql.DB) {
	// GET http://localhost:8000
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		response.OK(
			"Kasir Api, check documentation at "+r.Host+"/docs",
			nil,
			nil,
		).JSON(w, http.StatusOK)
	})

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

	CategoryRoute(mux, e, db)
	ProductRoute(mux, e, db)
	TransactionRoute(mux, e, db)
	ReportRoute(mux, e, db)
	// add other route...
}
