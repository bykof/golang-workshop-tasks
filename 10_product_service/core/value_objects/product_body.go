package value_objects

type ProductBody struct {
	Name     string  `form:"name" json:"name" binding:"required"`
	Price    float64 `form:"price" json:"price"`
	Category string  `form:"category" json:"category"`
}
