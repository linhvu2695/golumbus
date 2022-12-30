package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Author    string  `json:"author"`
	Price     float32 `json:"price"`
	SKU       string  `json:"sku"`
	CreatedOn string  `json:"-"`
	UpdatedOn string  `json:"-"`
	DeletedOn string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJson(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func (p *Product) FromJson(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, idx, err := findProduct(id)
	if err != nil {
		return err
	}
	
	p.ID = id
	productList[idx] = p
	return nil
}

func DeleteProduct(id int) error {
	_, idx, err := findProduct(id)
	if err != nil {
		return err
	}
	productList = append(productList[:idx], productList[idx+1:]...)
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for idx, prod := range productList {
		if prod.ID == id {
			return prod, idx, nil
		}
	}
	return nil, 0, ErrProductNotFound
}

func getNextId() int {
	pl := productList[len(productList)-1]
	return pl.ID + 1
}

var productList = []*Product{
	{
		ID:        1,
		Name:      "War and Peace",
		Author:    "Lev Tolstoy",
		Price:     140.99,
		SKU:       "abc123",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		ID:        2,
		Name:      "The Old Man and the Sea",
		Author:    "Ernest Hemingway",
		Price:     21.25,
		SKU:       "fjd334",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		ID:        3,
		Name:      "Frankenstein",
		Author:    "Marry Shelley",
		Price:     15.75,
		SKU:       "kbx898",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
