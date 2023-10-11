package storage

import "bootcamp-web/internal"

// NewStorageWarehouseDefaultValidator returns a new instance of StorageWarehouseDefaultValidator
func NewStorageWarehouseDefaultValidator(st StorageWarehouse) (s *StorageWarehouseDefaultValidator) {
	s = &StorageWarehouseDefaultValidator{
		st: st,
	}
	return
}

// StorageWarehouseDefaultValidator is an struct that represents the default validator of a warehouse storage
type StorageWarehouseDefaultValidator struct {
	// st is the storage of the warehouse
	st StorageWarehouse
}

// FindById returns the warehouse with the given Id
func (s *StorageWarehouseDefaultValidator) FindById(id int) (w internal.WarehouseDB, err error) {
	w, err = s.st.FindById(id)
	return
}

// FindByName returns the warehouse with the given name
func (s *StorageWarehouseDefaultValidator) FindByName(name string) (w internal.WarehouseDB, err error) {
	w, err = s.st.FindByName(name)
	return
}

// Add adds a warehouse to the storage
func (s *StorageWarehouseDefaultValidator) Add(w *internal.WarehouseDB) (err error) {
	// check if quantity is valid
	for _, qt := range w.Stock {
		if qt < 0 {
			err = ErrStorageWarehouseInvalidQuantity
			return
		}
	}

	// add warehouse
	err = s.st.Add(w)
	return
}

// Update updates the warehouse with the given Id
func (s *StorageWarehouseDefaultValidator) Update(w *internal.WarehouseDB) (err error) {
	// check if quantity is valid
	for _, qt := range w.Stock {
		if qt < 0 {
			err = ErrStorageWarehouseInvalidQuantity
			return
		}
	}

	// update warehouse
	err = s.st.Update(w)
	return
}