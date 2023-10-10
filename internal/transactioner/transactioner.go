package transactioner

import (
	"errors"
)

var (
	// ErrTransactionerWarehouseNotFound is the error returned when the warehouse is not found
	ErrTransactionerWarehouseNotFound = errors.New("warehouse not found")

	// ErrTransactionerOrderNotFound is the error returned when the order is not found
	ErrTransactionerOrderNotFound = errors.New("order not found")

	// ErrTransactionerWarehouseProductNotFound is the error returned when the product is not found in the warehouse
	ErrTransactionerWarehouseProductNotFound = errors.New("product not found in the warehouse")

	// ErrTransactionerOrderQuantityNotAvailable is the error returned when the quantity of a product in the order is not available in the warehouse
	ErrTransactionerOrderQuantityNotAvailable = errors.New("quantity of a product in the order is not available in the warehouse")
)

// Transactioner processes transactions between orders and warehouses
type Transactioner interface {
	// Fulfill processes the transaction between the order and the warehouse
	Fullfill(orderId, warehouseId int) (err error)
}