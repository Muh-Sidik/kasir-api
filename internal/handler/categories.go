package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Muh-Sidik/kasir-api/internal/model/dto"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/request"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/response"
	"github.com/Muh-Sidik/kasir-api/internal/service"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(srv service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: srv,
	}
}

// @Summary      Show categories
// @Description  get list category
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param		 name 		query		string 	false 	"Search by Category"
// @Param		 page		query		int	false	"Page number"
// @Param		 per_page	query		int	false	"Items per page"
// @Success      200  {object}  map[string]any
// @Router       /api/categories [get]
func (h *CategoryHandler) Categories(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	search := queryParam.Get("name")
	page := queryParam.Get("page")
	perPage := queryParam.Get("per_page")

	paginate := request.Paginate(page, perPage)
	category, total, err := h.service.GetCategories(paginate, search)

	if err != nil {
		response.Failed(
			"Failed get categories",
			err,
		).JSON(w, http.StatusInternalServerError)
		return
	}

	response.OK(
		"Successfully get categories",
		category,
		&response.Meta{
			Total: total,
			Page:  paginate.Page,
			Limit: paginate.Limit,
		},
	).JSON(w, http.StatusOK)
}

// @Summary      Create categories
// @Description  create a category
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param		 category	body		dto.CategoryRequest	true	"Add category"
// @Success      200  {object} 			map[string]any
// @Router       /api/categories [post]
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	body, err := request.BindJSON[dto.CategoryRequest](r)
	if err != nil {
		response.Failed(
			"Invalid Request",
			err,
		).JSON(w, http.StatusBadRequest)
		return
	}

	category, err := h.service.CreateCategory(&body)

	if err != nil {
		response.Failed(
			"Failed create category",
			err,
		).JSON(w, http.StatusInternalServerError)
		return
	}

	response.OK(
		"Successfully create category",
		category,
		nil,
	).JSON(w, http.StatusCreated)
}

// @Summary			Show a category
// @Description		get category by ID
// @Tags			Categories
// @Accept			json
// @Produce			json
// @Param			id	path		int		true	"Category ID"
// @Success			200	{object}	map[string]any
// @Router			/api/categories/{id} [get]
func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	category, err := h.service.GetCategoryByID(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Failed(
				"Not Found category",
				err,
			).JSON(w, http.StatusNotFound)
			return
		}

		response.Failed(
			"Failed get category",
			err,
		).JSON(w, http.StatusInternalServerError)
		return
	}

	response.OK(
		"Successfully get category",
		category,
		nil,
	).JSON(w, http.StatusOK)
}

// @Summary		Update a category
// @Description	Update category by ID
// @Tags			Categories
// @Accept			json
// @Produce		json
// @Param			id		path		string					true	"Category ID"
// @Param			account	body		model.Categories	true	"Update account"
// @Success		200		{object}	map[string]any
// @Router			/api/categories/{id} [put]
func (h *CategoryHandler) UpdateCategoryByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	body, err := request.BindJSON[dto.CategoryRequest](r)
	if err != nil {
		response.Failed(
			"Invalid Request",
			err,
		).JSON(w, http.StatusBadRequest)
		return
	}

	category, err := h.service.UpdateCategoryByID(id, &body)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Failed(
				"Not Found category",
				err,
			).JSON(w, http.StatusNotFound)
			return
		}

		response.Failed(
			"Failed update category",
			err,
		).JSON(w, http.StatusInternalServerError)
		return
	}

	response.OK(
		"Successfully update category",
		category,
		nil,
	).JSON(w, http.StatusOK)
}

// @Summary			Delete a category
// @Description		delete category by ID
// @Tags			Categories
// @Accept			json
// @Produce			json
// @Param			id	path		int		true	"Category ID"
// @Success			200	{object}	map[string]any
// @Router			/api/categories/{id} [delete]
func (h *CategoryHandler) DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.service.DeleteCategoryByID(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Failed(
				"Not Found category",
				err,
			).JSON(w, http.StatusNotFound)
			return
		}

		response.Failed(
			"Failed delete category",
			err,
		).JSON(w, http.StatusInternalServerError)
		return
	}

	response.OK(
		"Successfully delete category",
		nil,
		nil,
	).JSON(w, http.StatusOK)
}
