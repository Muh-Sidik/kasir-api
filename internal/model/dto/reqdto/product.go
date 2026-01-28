package reqdto

type ProductRequest struct {
	Name       string `json:"name" validate:"required,min=3"`
	Price      int    `json:"price" validate:"required,numeric"`
	Stock      int    `json:"stock" validate:"required,min=0,numeric"`
	CategoryID string `json:"category_id" validate:"required,uuid"`
}
