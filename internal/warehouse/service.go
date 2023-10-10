package warehouse

import (
	"bootcamp-web/internal/product"
	"errors"
)

var (
	// ErrServiceWarehouseNotFound is the error returned when the warehouse is not found
	ErrServiceWarehouseNotFound = errors.New("warehouse not found")

	// ErrServiceWarehouseProductNotFound is the error returned when the product is not found in the warehouse
	ErrServiceWarehouseProductNotFound = errors.New("product not found in the warehouse")
)

// WarehouseAttributes is a struct that represents the attributes of a warehouse
type WarehouseAttributes struct {
	// name is the name of the warehouse
	Name string

	// Stock is the stock of products in the warehouse with the quantity of each product
	Stock map[product.Product]int
}

// Warehouse is a struct that represents a warehouse to store products
type Warehouse struct {
	// Id is the unique identifier of the warehouse
	Id int

	// attributes is the attributes of the warehouse
	Attributes WarehouseAttributes
}

// ServiceWarehouse is the interface that wraps the basic methods of a warehouse service
type ServiceWarehouse interface {
	// FindById returns the warehouse with the given Id
	FindById(id int) (w Warehouse, err error)

	// FindByName returns the warehouse with the given name
	FindByName(name string) (w Warehouse, err error)

	// Add adds a warehouse to the storage
	Add(w *Warehouse) (err error)

	// Update updates the warehouse with the given Id
	Update(w *Warehouse) (err error)
}
