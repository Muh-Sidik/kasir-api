package handler

import (
	"github.com/Muh-Sidik/kasir-api/internal/pkg/validator"
	"github.com/Muh-Sidik/kasir-api/internal/service"
)

type Handler struct {
	ProductSrv  service.ProductService
	CategorySrv service.CategoryService
	Validator   validator.ValidatePkg
}
