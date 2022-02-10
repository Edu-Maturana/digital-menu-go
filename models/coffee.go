package models

type Coffee struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Img         string  `json:"img"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Coffees []Coffee
