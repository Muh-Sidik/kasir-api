package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Muh-Sidik/kasir-api/internal/model/dto"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/request"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/response"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/utils"
	"github.com/Muh-Sidik/kasir-api/internal/service"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(srv service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: srv,
	}
}

// @Summary      Show product
// @Description  get list product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param		 name 			query		string 	false 	"Search by Product Name"
// @Param		 categoryId 	query		string 	false 	"Filter by category id"
// @Param		 page			query		int		false	"Page number"
// @Param		 per_page		query		int		false	"Items per page"
// @Success      200  {object}  map[string]any
// @Router       /api/product [get]
func (h *ProductHandler) Products(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	page := queryParam.Get("page")
	perPage := queryParam.Get("per_page")
	paginate := request.Paginate(page, perPage)

	queryDto := &dto.ProductQuery{
		Name:       queryParam.Get("name"),
		CategoryID: queryParam.Get("categoryId"),
	}
	queryDto.Limit = paginate.Limit
	queryDto.Offset = paginate.Offset

	product, total, err := h.service.GetProduct(queryDto)

	if err != nil {
		response.Failed(
			"Failed get products",
			err,
		).JSON(w, http.StatusInternalServerError)
		return
	}

	response.OK(
		"Successfully get data products",
		product,
		&response.Meta{
			Total: total,
			Page:  paginate.Page,
			Limit: paginate.Limit,
		},
	).JSON(w, http.StatusOK)
}

// @Summary      Create product
// @Description  create a product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param		 product	body		dto.ProductRequest	true	"Add product"
// @Success      200  {object} 			map[string]any
// @Router       /api/product [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := request.BindJSON[dto.ProductRequest](r)
	if err != nil {
		response.Failed(
			"Invalid Request",
			err,
		).JSON(w, http.StatusBadRequest)
		return
	}

	product, err := h.service.CreateProduct(&body)

	if err != nil {
		if errors.Is(err, utils.ErrCategoryNotFound) {
			response.Failed(
				"Not Found category",
				err,
			).JSON(w, http.StatusNotFound)
			return
		}
		response.Failed(
			"Failed Create Product",
			err,
		).JSON(w, http.StatusInternalServerError)
		return
	}

	response.Created(
		"Successfully create product",
		product,
	).JSON(w, http.StatusCreated)
}

// @Summary			Show a product
// @Description		get product by ID
// @Tags			Product
// @Accept			json
// @Produce			json
// @Param			id	path		int		true	"Product ID"
// @Success			200	{object}	map[string]any
// @Router			/api/product/{id} [get]
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	product, err := h.service.GetProductByID(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Failed(
				"Not Found product",
				err,
			).JSON(w, http.StatusNotFound)
			return
		}

		response.Failed(
			"Failed get product",
			err,
		).JSON(w, http.StatusInternalServerError)
		return
	}

	response.OK(
		"Successfully get product",
		product,
		nil,
	).JSON(w, http.StatusOK)
}

// @Summary		Update a product
// @Description	Update product by ID
// @Tags			Product
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"Product ID"
// @Param			account	body		dto.ProductRequest	true	"Update product"
// @Success		200		{object}	map[string]any
// @Router			/api/product/{id} [put]
func (h *ProductHandler) UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	body, err := request.BindJSON[dto.ProductRequest](r)
	if err != nil {
		response.Failed(
			"Invalid Request",
			err,
		).JSON(w, http.StatusBadRequest)
		return
	}

	product, err := h.service.UpdateProductByID(id, &body)

	if err != nil {
		if errors.Is(err, utils.ErrCategoryNotFound) {
			response.Failed(
				"Not Found category",
				err,
			).JSON(w, http.StatusNotFound)
			return
		}

		if errors.Is(err, sql.ErrNoRows) {
			response.Failed(
				"Not Found product",
				err,
			).JSON(w, http.StatusNotFound)
			return
		}

		response.Failed(
			"Failed update product",
			err,
		).JSON(w, http.StatusInternalServerError)
		return
	}

	response.OK(
		"Successfully update product",
		product,
		nil,
	).JSON(w, http.StatusOK)
}

// @Summary			Delete a product
// @Description		delete product by ID
// @Tags			Product
// @Accept			json
// @Produce			json
// @Param			id	path		int		true	"Product ID"
// @Success			200	{object}	map[string]any
// @Router			/api/product/{id} [delete]
func (h *ProductHandler) DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.service.DeleteProductByID(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Failed(
				"Not Found product",
				err,
			).JSON(w, http.StatusNotFound)
			return
		}

		response.Failed(
			"Failed delete product",
			err,
		).JSON(w, http.StatusInternalServerError)
		return
	}

	response.OK(
		"Successfully delete product",
		nil,
		nil,
	).JSON(w, http.StatusOK)
}
