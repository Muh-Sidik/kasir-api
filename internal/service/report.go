package service

import (
	"github.com/Muh-Sidik/kasir-api/internal/model"
	"github.com/Muh-Sidik/kasir-api/internal/model/dto"
	"github.com/Muh-Sidik/kasir-api/internal/repository"
)

type ReportService interface {
	GetReport(param *dto.ReportParam) (*model.TopProductReport, error)
}

type reportService struct {
	reportRepo repository.ReportRepository
}

func NewReportService(reportRepo repository.ReportRepository) ReportService {
	return &reportService{
		reportRepo: reportRepo,
	}
}

func (s *reportService) GetReport(param *dto.ReportParam) (*model.TopProductReport, error) {
	if _, _, err := param.ParseDates(); err != nil {
		return nil, err
	}

	report, err := s.reportRepo.Report(param)
	if err != nil {
		return nil, err
	}

	return report, nil
}
