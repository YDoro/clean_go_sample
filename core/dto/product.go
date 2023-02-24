package dto

import (
	"encoding/json"
	"io"
)

// CreateProductRequest is a representation of the request body to create product
type CreateProductRequest struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

func FromJSONCreateProductRequest(body io.Reader) (*CreateProductRequest, error) {
	cpr := CreateProductRequest{}
	if err := json.NewDecoder(body).Decode(&cpr); err != nil {
		return nil, err
	}

	return &cpr, nil
}
