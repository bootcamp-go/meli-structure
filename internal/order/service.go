package order

import (
	"bootcamp-web/internal/product"
	"errors"
)

var (
	// ErrServiceOrderNotFound is the error returned when the order is not found
	ErrServiceOrderNotFound = errors.New("order not found")

	// ErrServiceOrderProductNotFound is the error returned when the product is not found in the order
	ErrServiceOrderProductNotFound = errors.New("product not found in the order")
)

// Order is a struct that represents an order
type Order struct {
	// Id is the unique identifier of the order
	Id int

	// Products is the list of products in the order
	// - key: product
	// - value: quantity
	Products map[product.Product]int
}

// ServiceOrder is the interface that wraps the basic methods of an order service
type ServiceOrder interface {
	// FindById returns the order with the given Id
	FindById(id int) (o Order, err error)

	// Add adds an order to the storage
	Add(o *Order) (err error)

	// Update updates the order with the given Id
	Update(o *Order) (err error)
}