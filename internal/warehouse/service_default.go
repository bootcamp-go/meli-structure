package warehouse

import "bootcamp-web/internal/product"

// ServiceWarehouseDefault is the default implementation of ServiceWarehouse
type ServiceWarehouseDefault struct {
	// storage is the storage of warehouses
	storage StorageWarehouse

	// catalog is the catalog of products
	catalog product.CatalogProduct
}

// NewServiceWarehouseDefault returns a new instance of ServiceWarehouseDefault
func NewServiceWarehouseDefault(storage StorageWarehouse, catalog product.CatalogProduct) (s *ServiceWarehouseDefault) {
	s = &ServiceWarehouseDefault{
		storage: storage,
		catalog: catalog,
	}
	return
}

// FindById returns the warehouse with the given Id
func (s *ServiceWarehouseDefault) FindById(id int) (w Warehouse, err error) {
	// get the warehouse from the storage
	wh, err := s.storage.FindById(id)
	if err != nil {
		err = ErrServiceWarehouseNotFound
		return
	}

	// serialize the warehouse
	w.Id = wh.Id
	w.Attributes.Name = wh.Name
	w.Attributes.Stock = make(map[product.Product]int)
	for pr, qt := range wh.Stock {
		var p product.Product
		p, err = s.catalog.FindProductByName(pr)
		if err != nil {
			err = ErrServiceWarehouseProductNotFound
			return
		}

		w.Attributes.Stock[p] = qt
	}

	return
}

// FindByName returns the warehouse with the given name
func (s *ServiceWarehouseDefault) FindByName(name string) (w Warehouse, err error) {
	// get the warehouse from the storage
	wh, err := s.storage.FindByName(name)
	if err != nil {
		err = ErrServiceWarehouseNotFound
		return
	}

	// serialize the warehouse
	w.Id = wh.Id
	w.Attributes.Name = wh.Name
	w.Attributes.Stock = make(map[product.Product]int)
	for pr, qt := range wh.Stock {
		var p product.Product
		p, err = s.catalog.FindProductByName(pr)
		if err != nil {
			err = ErrServiceWarehouseProductNotFound
			return
		}

		w.Attributes.Stock[p] = qt
	}

	return
}

// Add adds a warehouse to the storage
func (s *ServiceWarehouseDefault) Add(w *Warehouse) (err error) {
	// serialize the warehouse
	wh := WarehouseDB{
		Id:    w.Id,
		Name:  w.Attributes.Name,
		Stock: make(map[string]int, 0),
	}
	for p, q := range w.Attributes.Stock {
		wh.Stock[p.Name()] = q
	}

	// add the warehouse to the storage
	err = s.storage.Add(&wh)
	if err != nil {
		return
	}

	return
}

// Update updates the warehouse with the given Id
func (s *ServiceWarehouseDefault) Update(w *Warehouse) (err error) {
	// serialize the warehouse
	wh := WarehouseDB{
		Id:    w.Id,
		Name:  w.Attributes.Name,
		Stock: make(map[string]int, 0),
	}
	for p, q := range w.Attributes.Stock {
		wh.Stock[p.Name()] = q
	}

	// update or create the warehouse in the storage
	err = s.storage.Update(&wh)
	if err != nil {
		return
	}

	return
}