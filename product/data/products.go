package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int `json: "id"`
	Name        string `json: "name"`
	Description string 	`json: "description`
	Price       int     `josn: "Price"`
	SKU         string  `json: "sku"`
	CreatedOn   string  `json: "_"` 
	DeletedOn   string  `json: "_"`
}

type Products []*Product

	func (p*Products)ToJSON(w io.Writer)error{
			e := json.NewEncoder(w)
				return e.Encode(p)
	}
	

func GetProducts()Products {
	return productList

}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "tamil",
		Description: "milk coffe",
		Price:       250,
		SKU:         "abcdes",
		CreatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "tamil",
		Description: "milk tea",
		Price:       200,
		SKU:         "abcdefgh",
		CreatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}
