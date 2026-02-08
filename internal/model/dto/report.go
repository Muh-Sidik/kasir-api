package dto

import (
	"fmt"
	"time"
)

type ReportParam struct {
	StartDate string
	EndDate   string
}

func (p *ReportParam) ParseDates() (time.Time, time.Time, error) {
	var startDate, endDate time.Time
	var err error

	if p.StartDate != "" {
		startDate, err = time.Parse("2006-01-02", p.StartDate)
		if err != nil {
			return time.Time{}, time.Time{}, fmt.Errorf("invalid start_date format: %w", err)
		}
	} else {
		startDate = time.Now().Truncate(24 * time.Hour)
	}

	if p.EndDate != "" {
		endDate, err = time.Parse("2006-01-02", p.EndDate)
		if err != nil {
			return time.Time{}, time.Time{}, fmt.Errorf("invalid end_date format: %w", err)
		}
	} else {
		endDate = time.Now().Truncate(24 * time.Hour)
	}

	// Validasi: start_date tidak boleh setelah end_date
	if startDate.After(endDate) {
		return time.Time{}, time.Time{}, fmt.Errorf("start_date cannot be after end_date")
	}

	return startDate, endDate, nil
}
