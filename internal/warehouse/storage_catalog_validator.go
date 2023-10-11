package warehouse

import "bootcamp-web/internal/product"

// NewStorageWarehouseCatalogValidator returns a new instance of StorageWarehouseCatalogValidator
func NewStorageWarehouseCatalogValidator(st StorageWarehouse, ct product.CatalogProduct) (s *StorageWarehouseCatalogValidator) {
	s = &StorageWarehouseCatalogValidator{
		st: st,
		ct: ct,
	}
	return
}

// StorageWarehouseCatalogValidator is an struct that represents the catalog validator of a warehouse storage
type StorageWarehouseCatalogValidator struct {
	// st is the storage of the warehouse
	st StorageWarehouse

	// ct is the catalog of products
	ct product.CatalogProduct
}

// FindById returns the warehouse with the given Id
func (s *StorageWarehouseCatalogValidator) FindById(id int) (w WarehouseDB, err error) {
	w, err = s.st.FindById(id)
	return
}

// FindByName returns the warehouse with the given name
func (s *StorageWarehouseCatalogValidator) FindByName(name string) (w WarehouseDB, err error) {
	w, err = s.st.FindByName(name)
	return
}

// Add adds a warehouse to the storage
func (s *StorageWarehouseCatalogValidator) Add(w *WarehouseDB) (err error) {
	// check if products are in catalog
	for pr := range w.Stock {
		_, err = s.ct.FindProductByName(pr)
		if err != nil {
			err = ErrStorageWarehouseProductNotFound
			return
		}
	}

	// add warehouse
	err = s.st.Add(w)
	return
}

// Update updates the warehouse with the given Id
func (s *StorageWarehouseCatalogValidator) Update(w *WarehouseDB) (err error) {
	// check if products are in catalog
	for pr := range w.Stock {
		_, err = s.ct.FindProductByName(pr)
		if err != nil {
			err = ErrStorageWarehouseProductNotFound
			return
		}
	}

	// update warehouse
	err = s.st.Update(w)
	return
}