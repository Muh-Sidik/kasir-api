package dto

import "github.com/Muh-Sidik/kasir-api/internal/pkg/request"

type ProductQuery struct {
	Name       string `json:"name" validate:"required,min=3"`
	CategoryID string `json:"category_id" validate:"required,uuid"`
	request.PaginateQuery
}

type ProductRequest struct {
	Name       string `json:"name" validate:"required,min=3"`
	Price      int    `json:"price" validate:"required,numeric"`
	Stock      int    `json:"stock" validate:"required,min=0,numeric"`
	CategoryID string `json:"category_id" validate:"required,uuid"`
}
