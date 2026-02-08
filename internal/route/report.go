package route

import (
	"database/sql"
	"net/http"

	"github.com/Muh-Sidik/kasir-api/config"
	"github.com/Muh-Sidik/kasir-api/internal/handler"
	"github.com/Muh-Sidik/kasir-api/internal/repository"
	"github.com/Muh-Sidik/kasir-api/internal/service"
)

func ReportRoute(mux *http.ServeMux, e *config.Env, db *sql.DB) {
	handler := handler.NewReportHandler(
		service.NewReportService(
			repository.NewReportRepository(db),
		),
	)

	// GET http://localhost:8000/api/report
	mux.HandleFunc("GET /api/report", handler.Report)
}
