package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Muh-Sidik/kasir-api/internal/model"
	"github.com/Muh-Sidik/kasir-api/internal/model/dto"
)

type ReportRepository interface {
	Report(param *dto.ReportParam) (*model.TopProductReport, error)
}

type reportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return &reportRepository{
		db: db,
	}
}

func (r *reportRepository) Report(param *dto.ReportParam) (*model.TopProductReport, error) {
	query := `
		WITH report_data AS (
			SELECT 
				COALESCE(SUM(td.subtotal), 0)::BIGINT as total_revenue,
				COALESCE(COUNT(DISTINCT t.id), 0)::BIGINT as total_transaction,
				p.name as top_product_name,
				COALESCE(SUM(td.quantity), 0)::BIGINT as product_quantity,
				ROW_NUMBER() OVER (ORDER BY SUM(td.quantity) DESC) as rn
			FROM transactions t
			JOIN transaction_details td ON t.id = td.transaction_id
			JOIN product p ON td.product_id = p.id
			WHERE t.created_at >= $1::date 
			  AND t.created_at < ($2::date + INTERVAL '1 day')
			GROUP BY p.id, p.name
		)
		SELECT 
			total_revenue,
			total_transaction,
			top_product_name,
			product_quantity
		FROM report_data
		WHERE rn = 1
	`

	// Handle default date (today jika kosong)
	startDate := param.StartDate
	endDate := param.EndDate

	if startDate == "" {
		startDate = time.Now().Format("2006-01-02")
	}
	if endDate == "" {
		endDate = time.Now().Format("2006-01-02")
	}

	var result model.TopProductReport
	err := r.db.QueryRow(query, startDate, endDate).Scan(
		&result.TotalRevenue,
		&result.TotalTransaction,
		&result.Products.ProductName,
		&result.Products.QuantitySold,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// Return empty report jika tidak ada data
			return &model.TopProductReport{
				RevenueReport:     model.RevenueReport{TotalRevenue: 0},
				TransactionReport: model.TransactionReport{TotalTransaction: 0},
				Products: model.TopProduct{
					ProductName:  "",
					QuantitySold: 0,
				},
			}, nil
		}
		return nil, fmt.Errorf("query report failed: %w", err)
	}

	return &result, nil
}
