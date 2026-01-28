package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Muh-Sidik/kasir-api/internal/model/dto/reqdto"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/request"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/response"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/utils"
)

// @Summary      Show product
// @Description  get list product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param		 page		query		int	false	"Page number"
// @Param		 per_page	query		int	false	"Items per page"
// @Success      200  {object}  map[string]any
// @Router       /api/product [get]
func (h *Handler) Products(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	perPage := r.URL.Query().Get("per_page")

	paginate := request.Paginate(page, perPage)
	product, total, err := h.ProductSrv.GetProduct(paginate)

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
// @Param		 product	body		reqdto.ProductRequest	true	"Add product"
// @Success      200  {object} 			map[string]any
// @Router       /api/product [post]
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := request.BindJSON[reqdto.ProductRequest](r)
	if err != nil {
		response.Failed(
			"Invalid Request",
			err,
		).JSON(w, http.StatusBadRequest)
		return
	}

	product, err := h.ProductSrv.CreateProduct(&body)

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
func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	product, err := h.ProductSrv.GetProductByID(id)

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
// @Param			account	body		reqdto.ProductRequest	true	"Update product"
// @Success		200		{object}	map[string]any
// @Router			/api/product/{id} [put]
func (h *Handler) UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	body, err := request.BindJSON[reqdto.ProductRequest](r)
	if err != nil {
		response.Failed(
			"Invalid Request",
			err,
		).JSON(w, http.StatusBadRequest)
		return
	}

	product, err := h.ProductSrv.UpdateProductByID(id, &body)

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
func (h *Handler) DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.ProductSrv.DeleteProductByID(id)

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
