package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Muh-Sidik/kasir-api/internal/model"
)

// @Summary      Show categories
// @Description  get list category
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]any
// @Router       /api/categories [get]
func (h *Handler) Categories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	product, err := h.getCategory()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(map[string]string{
			"status":  "FAILED",
			"message": "failed get categories",
			"error":   err.Error(),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]any{
		"status": "OK",
		"data":   product,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary      Create categories
// @Description  create a category
// @Tags         Categories
// @Accept       json
// @Produce      json
//	@Param		 category	body		model.Categories	true	"Add category"
// @Success      200  {object} 			map[string]any
// @Router       /api/categories [post]

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body := new(model.Categories)
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(map[string]string{
			"status":  "FAILED",
			"message": "invalid request",
			"error":   err.Error(),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	product, err := h.createCategory(body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(map[string]string{
			"status":  "FAILED",
			"message": "failed create category",
			"error":   err.Error(),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]any{
		"status": "OK",
		"data":   product,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(map[string]string{
			"status":  "FAILED",
			"message": "invalid request",
			"error":   err.Error(),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	product, err := h.getCategoryByID(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(map[string]string{
				"status":  "FAILED",
				"message": "Not Found category",
				"error":   err.Error(),
			})
		}

		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(map[string]string{
			"status":  "FAILED",
			"message": "failed get category",
			"error":   err.Error(),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]any{
		"status": "OK",
		"data":   product,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(map[string]string{
			"status":  "FAILED",
			"message": "invalid request",
			"error":   err.Error(),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	body := new(model.Categories)
	err = json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(map[string]string{
			"status":  "FAILED",
			"message": "invalid request",
			"error":   err.Error(),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	product, err := h.updateCategoryByID(id, body)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(map[string]string{
				"status":  "FAILED",
				"message": "Not Found category",
				"error":   err.Error(),
			})
		}

		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(map[string]string{
			"status":  "FAILED",
			"message": "failed update category",
			"error":   err.Error(),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]any{
		"status": "OK",
		"data":   product,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(map[string]string{
			"status":  "FAILED",
			"message": "invalid request",
			"error":   err.Error(),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err = h.deleteCategoryByID(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(map[string]string{
				"status":  "FAILED",
				"message": "Not Found category",
				"error":   err.Error(),
			})
		}

		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(map[string]string{
			"status":  "FAILED",
			"message": "failed delete category",
			"error":   err.Error(),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]any{
		"status": "OK",
		"data":   nil,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) getCategory() ([]*model.Categories, error) {
	query := `SELECT id, name, description FROM categories WHERE 1=1`

	rows, err := h.DB.Query(query)

	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	defer rows.Close()

	categories := make([]*model.Categories, 0)

	for rows.Next() {
		var category model.Categories
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	return categories, nil
}

func (h *Handler) createCategory(body *model.Categories) (*model.Categories, error) {
	rows := h.DB.QueryRow(
		`INSERT INTO categories(name, description) VALUES(?,?) RETURNING id`,
		body.Name,
		body.Description,
	)

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var category model.Categories
	if err := rows.Scan(&category.ID); err != nil {
		return nil, err
	}

	category.Name = body.Name
	category.Description = body.Description

	return &category, nil
}

func (h *Handler) getCategoryByID(id int) (*model.Categories, error) {
	query := `SELECT id, name, description FROM categories WHERE id = ?`

	rows := h.DB.QueryRow(query, id)

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var category model.Categories
	err := rows.Scan(
		&category.ID,
		&category.Name,
		&category.Description,
	)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (h *Handler) deleteCategoryByID(id int) error {
	_, err := h.DB.Exec(
		`DELETE FROM categories WHERE id = ?`,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) updateCategoryByID(id int, body *model.Categories) (*model.Categories, error) {
	rows := h.DB.QueryRow(
		`UPDATE categories SET name = ?, description = ? WHERE id = ? RETURNING id`,
		body.Name,
		body.Description,
		id,
	)

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var category model.Categories
	if err := rows.Scan(&category.ID); err != nil {
		return nil, err
	}

	category.Name = body.Name
	category.Description = body.Description

	return &category, nil
}
