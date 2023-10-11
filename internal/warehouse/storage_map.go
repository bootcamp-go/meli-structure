package warehouse

import "bootcamp-web/internal/product"

type StorageWarehouseMap struct {
	// db is the database of warehouses (key: id, value: warehouse)
	db map[int]WarehouseDB

	// lastId is the next id to be assigned to a warehouse
	lastId int

	// catalog is the catalog of products
	catalog product.CatalogProduct
}

// NewStorageWarehouseMap returns a new instance of StorageWarehouseMap
func NewStorageWarehouseMap(db map[int]WarehouseDB, lastId int, catalog product.CatalogProduct) (s *StorageWarehouseMap) {
	s = &StorageWarehouseMap{
		db:     db,
		lastId: lastId,
		catalog: catalog,
	}
	return
}

// FindById returns the warehouse with the given Id
func (s *StorageWarehouseMap) FindById(id int) (w WarehouseDB, err error) {
	// get warehouse
	w, ok := s.db[id]
	if !ok {
		err = ErrStorageWarehouseNotFound
		return
	}

	return
}

// FindByName returns the warehouse with the given name
func (s *StorageWarehouseMap) FindByName(name string) (w WarehouseDB, err error) {
	// check if warehouse exists
	var exists bool
	var id int
	for ix, wh := range s.db {
		if wh.Name == name {
			exists = true
			id = ix
		}
	}
	if !exists {
		err = ErrStorageWarehouseNotFound
		return
	}

	// get warehouse
	w = s.db[id]
	return
}

// Add adds a warehouse to the storage
func (s *StorageWarehouseMap) Add(w *WarehouseDB) (err error) {
	// validation (should be decoupled)
	// check if quantity is valid
	for _, qt := range w.Stock {
		if qt < 0 {
			err = ErrStorageWarehouseInvalidQuantity
			return
		}
	}
	// check if products are in catalog
	for pr := range w.Stock {
		_, err = s.catalog.FindProductByName(pr)
		if err != nil {
			err = ErrStorageWarehouseProductNotFound
			return
		}
	}

	// set id
	s.lastId++
	w.Id = s.lastId

	// add warehouse
	s.db[s.lastId] = *w
	return
}

// Update updates the warehouse with the given Id
func (s *StorageWarehouseMap) Update(w *WarehouseDB) (err error) {
	// validation (should be decoupled)
	// check if quantity is valid
	for _, qt := range w.Stock {
		if qt < 0 {
			err = ErrStorageWarehouseInvalidQuantity
			return
		}
	}
	// check if products are in catalog
	for pr := range w.Stock {
		_, err = s.catalog.FindProductByName(pr)
		if err != nil {
			err = ErrStorageWarehouseProductNotFound
			return
		}
	}

	// check if warehouse exists
	var exists bool
	var id int
	for ix := range s.db {
		if ix == w.Id {
			exists = true
			id = ix
		}
	}
	if !exists {
		err = ErrStorageWarehouseNotFound
		return
	}

	// validate warehouse
	for pr := range w.Stock {
		_, err = s.catalog.FindProductByName(pr)
		if err != nil {
			return
		}
	}

	// update warehouse
	s.db[id] = *w
	return
}