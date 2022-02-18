package models

type Coffee struct {
	Id          string  `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Img         string  `json:"img" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

type Coffees []Coffee
