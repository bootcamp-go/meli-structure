package storage

import (
	"bootcamp-web/internal"
	"errors"
)

var (
	// ErrStorageWarehouseNotFound is an error that occurs when a warehouse is not found in the storage
	ErrStorageWarehouseNotFound = errors.New("storage-warehouse: warehouse not found")

	// ErrStorageWarehouseProductNotFound is an error that occurs when a product is not found in the warehouse
	ErrStorageWarehouseProductNotFound = errors.New("storage-warehouse: product not found in the warehouse")

	// ErrStorageWarehouseInvalidQuantity is an error that occurs when the quantity of a product is invalid
	ErrStorageWarehouseInvalidQuantity = errors.New("storage-warehouse: invalid quantity")
)

// StorageWarehouse is an interface that represents a storage for warehouses
type StorageWarehouse interface {
	// FindById returns the warehouse with the given Id
	FindById(id int) (w internal.WarehouseDB, err error)

	// FindByName returns the warehouse with the given name
	FindByName(name string) (w internal.WarehouseDB, err error)

	// Add adds a warehouse to the storage
	Add(w *internal.WarehouseDB) (err error)

	// Update updates the warehouse with the given Id
	Update(w *internal.WarehouseDB) (err error)
}