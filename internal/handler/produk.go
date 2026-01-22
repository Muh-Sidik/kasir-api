package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Muh-Sidik/kasir-api/internal/model"
)

type Handler struct {
	DB *sql.DB
}

// @Summary      Show product
// @Description  get list product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]any
// @Router       /api/product [get]
func (h *Handler) Products(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	product, err := h.getProduct()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(map[string]string{
			"status":  "FAILED",
			"message": "failed get products",
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

// @Summary      Create product
// @Description  create a product
// @Tags         Product
// @Accept       json
// @Produce      json
//
//	@Param		 product	body		model.Produk	true	"Add product"
//
// @Success      200  {object} 			map[string]any
// @Router       /api/product [post]
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body := new(model.Produk)
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

	product, err := h.createProduct(body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(map[string]string{
			"status":  "FAILED",
			"message": "failed create product",
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

// @Summary			Show a product
// @Description		get product by ID
// @Tags			Product
// @Accept			json
// @Produce			json
// @Param			id	path		int		true	"Product ID"
// @Success			200	{object}	map[string]any
// @Router			/api/product/{id} [get]
func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
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
	product, err := h.getProductByID(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(map[string]string{
				"status":  "FAILED",
				"message": "Not Found product",
				"error":   err.Error(),
			})
		}

		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(map[string]string{
			"status":  "FAILED",
			"message": "failed get product",
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

// @Summary		Update a product
// @Description	Update product by ID
// @Tags			Product
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"Product ID"
// @Param			account	body		model.Produk	true	"Update product"
// @Success		200		{object}	map[string]any
// @Router			/api/product/{id} [put]
func (h *Handler) UpdateProductByID(w http.ResponseWriter, r *http.Request) {
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

	body := new(model.Produk)
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

	product, err := h.updateProductByID(id, body)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(map[string]string{
				"status":  "FAILED",
				"message": "Not Found product",
				"error":   err.Error(),
			})
		}

		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(map[string]string{
			"status":  "FAILED",
			"message": "failed update product",
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

// @Summary			Delete a product
// @Description		delete product by ID
// @Tags			Product
// @Accept			json
// @Produce			json
// @Param			id	path		int		true	"Product ID"
// @Success			200	{object}	map[string]any
// @Router			/api/product/{id} [delete]
func (h *Handler) DeleteProductByID(w http.ResponseWriter, r *http.Request) {
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

	err = h.deleteProductByID(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(map[string]string{
				"status":  "FAILED",
				"message": "Not Found product",
				"error":   err.Error(),
			})
		}

		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(map[string]string{
			"status":  "FAILED",
			"message": "failed delete product",
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

func (h *Handler) getProduct() ([]*model.Produk, error) {
	query := `SELECT id, nama, harga, stok FROM produk WHERE 1=1`

	rows, err := h.DB.Query(query)

	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	defer rows.Close()

	listProduct := make([]*model.Produk, 0)

	for rows.Next() {
		var product model.Produk
		err := rows.Scan(
			&product.ID,
			&product.Nama,
			&product.Harga,
			&product.Stok,
		)

		if err != nil {
			return nil, err
		}

		listProduct = append(listProduct, &product)
	}

	return listProduct, nil
}

func (h *Handler) createProduct(body *model.Produk) (*model.Produk, error) {
	rows := h.DB.QueryRow(
		`INSERT INTO produk(nama, harga, stok) VALUES(?,?,?) RETURNING id`,
		body.Nama,
		body.Harga,
		body.Stok,
	)

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var product model.Produk
	if err := rows.Scan(&product.ID); err != nil {
		return nil, err
	}

	product.Nama = body.Nama
	product.Harga = body.Harga
	product.Stok = body.Stok

	return &product, nil
}

func (h *Handler) getProductByID(id int) (*model.Produk, error) {
	query := `SELECT id, nama, harga, stok FROM produk WHERE id = ?`

	rows := h.DB.QueryRow(query, id)

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var product model.Produk
	err := rows.Scan(
		&product.ID,
		&product.Nama,
		&product.Harga,
		&product.Stok,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (h *Handler) deleteProductByID(id int) error {
	_, err := h.DB.Exec(
		`DELETE FROM produk WHERE id = ?`,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) updateProductByID(id int, body *model.Produk) (*model.Produk, error) {
	rows := h.DB.QueryRow(
		`UPDATE produk SET nama = ?, harga = ?, stok = ? WHERE id = ? RETURNING id`,
		body.Nama,
		body.Harga,
		body.Stok,
		id,
	)

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var product model.Produk
	if err := rows.Scan(&product.ID); err != nil {
		return nil, err
	}

	product.Nama = body.Nama
	product.Harga = body.Harga
	product.Stok = body.Stok

	return &product, nil
}
