package handler

import (
	"net/http"

	"github.com/Muh-Sidik/kasir-api/internal/model/dto"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/response"
	"github.com/Muh-Sidik/kasir-api/internal/service"
)

type ReportHandler struct {
	reportService service.ReportService
}

func NewReportHandler(reportService service.ReportService) *ReportHandler {
	return &ReportHandler{
		reportService: reportService,
	}
}

// @Summary      Show report
// @Description  get report
// @Tags         Report
// @Accept       json
// @Produce      json
// @Param		 start_date 			query		string 	false 	"Start Date"
// @Param		 end_date 				query		string 	false 	"End Date"
// @Success      200  {object}  map[string]any
// @Router       /api/report [get]
func (h *ReportHandler) Report(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	start := queryParam.Get("start_date")
	end := queryParam.Get("end_date")

	queryDto := &dto.ReportParam{
		StartDate: start,
		EndDate:   end,
	}

	topProduct, err := h.reportService.GetReport(queryDto)

	if err != nil {
		response.Failed(
			"Failed get top product",
			err,
		).JSON(w, http.StatusInternalServerError)
		return
	}

	response.OK(
		"Successfully get data top product",
		topProduct,
		nil,
	).JSON(w, http.StatusOK)
}
