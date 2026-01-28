package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/Muh-Sidik/kasir-api/internal/model/dto/reqdto"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/request"
	"github.com/Muh-Sidik/kasir-api/internal/pkg/response"
)

// @Summary      Show categories
// @Description  get list category
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param		 page		query		int	false	"Page number"
// @Param		 per_page	query		int	false	"Items per page"
// @Success      200  {object}  map[string]any
// @Router       /api/categories [get]
func (h *Handler) Categories(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	perPage := r.URL.Query().Get("per_page")

	paginate := request.Paginate(page, perPage)
	category, total, err := h.CategorySrv.GetCategories(paginate)

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
// @Param		 category	body		reqdto.CategoryRequest	true	"Add category"
// @Success      200  {object} 			map[string]any
// @Router       /api/categories [post]
func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	body, err := request.BindJSON[reqdto.CategoryRequest](r)
	if err != nil {
		response.Failed(
			"Invalid Request",
			err,
		).JSON(w, http.StatusBadRequest)
		return
	}

	category, err := h.CategorySrv.CreateCategory(&body)

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
func (h *Handler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	idParams := r.PathValue("id")

	id, err := strconv.Atoi(idParams)
	if err != nil {
		response.Failed(
			"Invalid Request",
			err,
		).JSON(w, http.StatusBadRequest)
		return
	}
	category, err := h.CategorySrv.GetCategoryByID(id)

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
// @Param			id		path		int					true	"Category ID"
// @Param			account	body		model.Categories	true	"Update account"
// @Success		200		{object}	map[string]any
// @Router			/api/categories/{id} [put]
func (h *Handler) UpdateCategoryByID(w http.ResponseWriter, r *http.Request) {
	idParams := r.PathValue("id")

	id, err := strconv.Atoi(idParams)
	if err != nil {
		response.Failed(
			"Invalid Request",
			err,
		).JSON(w, http.StatusBadRequest)
		return
	}

	body, err := request.BindJSON[reqdto.CategoryRequest](r)
	if err != nil {
		response.Failed(
			"Invalid Request",
			err,
		).JSON(w, http.StatusBadRequest)
		return
	}

	category, err := h.CategorySrv.UpdateCategoryByID(id, &body)

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
func (h *Handler) DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	idParams := r.PathValue("id")

	id, err := strconv.Atoi(idParams)
	if err != nil {
		response.Failed(
			"Invalid Request",
			err,
		).JSON(w, http.StatusBadRequest)
		return
	}

	err = h.CategorySrv.DeleteCategoryByID(id)

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
