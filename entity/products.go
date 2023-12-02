package entity

type UpdateProductInput struct {
	Name  string
	Price float64
	Stock int
}

type Products struct {
	Product_id int
	Name       string
	Price      float64
	Stock      int
}
