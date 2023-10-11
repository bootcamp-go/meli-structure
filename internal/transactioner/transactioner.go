package transactioner

import (
	"errors"
)

var (
	// Order errors
	// ErrTransactionerOrderQuantityNotPositive is the error returned when the quantity of a product in the order is not positive
	ErrTransactionerOrderQuantityNotPositive = errors.New("quantity of a product in the order is not positive")
	// ErrTransactionerOrderQuantityNotAvailable is the error returned when the quantity of a product in the order is not available in the warehouse
	ErrTransactionerOrderQuantityNotAvailable = errors.New("quantity of a product in the order is not available in the warehouse")

	// Warehouse errors
	// ErrTransactionerWarehouseNotFound is the error returned when the warehouse is not found
	ErrTransactionerWarehouseNotFound = errors.New("warehouse not found")
	// ErrTransactionerWarehouseProductNotFound is the error returned when the product is not found in the warehouse
	ErrTransactionerWarehouseProductNotFound = errors.New("product not found in the warehouse")
	// ErrTransactionerWarehouseProductQuantityInvalid is the error returned when the quantity of a product is invalid
	ErrTransactionerWarehouseProductQuantityInvalid = errors.New("invalid quantity of a product in the warehouse")

)

type Order struct {
	// Name is the name of the order
	Name string
	// Products is the list of products in the order with the quantity of each product
	// - key: product name
	// - value: quantity
	Products map[string]int
}
	

// Transactioner processes transactions between orders and warehouses
type Transactioner interface {
	// Fulfill processes the transaction between the order and the warehouse
	Fullfill(order Order, warehouseName string) (err error)
}