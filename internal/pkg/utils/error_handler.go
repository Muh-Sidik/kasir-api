package utils

import (
	"net/http"

	"github.com/Muh-Sidik/kasir-api/internal/pkg/response"
)

func CustomHandler(m string, code int, err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Failed(m, err).JSON(w, code)
	}
}
