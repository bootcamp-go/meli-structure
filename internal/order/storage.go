package order

import "errors"

var (
	// ErrStorageOrderNotFound is the error returned when the order is not found
	ErrStorageOrderNotFound = errors.New("order not found")
)

// OrderDB is a struct that represents an order
type OrderDB struct {
	// Id is the unique identifier of the order
	Id int

	// Products is the list of products in the order
	// - key: product name
	// - value: quantity
	Products map[string]int
}

// StorageOrder is the interface that wraps the basic methods of an order storage
type StorageOrder interface {
	// FindById returns the order with the given Id
	FindById(id int) (o OrderDB, err error)

	// Add adds an order to the storage
	Add(o *OrderDB) (err error)

	// Update updates the order with the given Id
	Update(o *OrderDB) (err error)
}