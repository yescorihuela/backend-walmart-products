package response

type ProductResponse struct {
	Id          uint    `json:"id"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Image       string  `json:"imagen"`
	Price       float32 `json:"price"`
	HasDiscount bool    `json:"has_discount"`
	Discount    uint    `json:"discount,omitempty"`
}
