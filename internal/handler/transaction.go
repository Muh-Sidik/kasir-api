package handler

import (
	"net/http"

	"github.com/Muh-Sidik/kasir-api/internal/model/dto"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/request"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/response"
	"github.com/Muh-Sidik/kasir-api/internal/service"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(srv service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: srv,
	}
}

// @Summary      Create Checkout
// @Description  create checkout
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Param		 checkout	body		dto.CheckoutRequest	true	"Add checkout"
// @Success      200  {object} 			map[string]any
// @Router       /api/checkout [post]
func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	body, err := request.BindJSON[dto.CheckoutRequest](r)

	if err != nil {
		response.Failed(
			"Invalid Request",
			err,
		).JSON(w, http.StatusBadRequest)
		return
	}

	transaction, err := h.service.CreateCheckout(body.Items)

	if err != nil {
		response.Failed(
			"Failed Create Checkout",
			err,
		).JSON(w, http.StatusInternalServerError)
		return
	}

	response.Created(
		"Successfully create checkout",
		transaction,
	).JSON(w, http.StatusCreated)
}
