package order

import "bootcamp-web/internal/product"

// ServiceOrderDefault is the default implementation of ServiceOrder
type ServiceOrderDefault struct {
	// storage is the storage to manage orders
	storage StorageOrder

	// catalog is the catalog of products
	catalog product.CatalogProduct
}

// NewServiceOrderDefault returns a new instance of ServiceOrderDefault
func NewServiceOrderDefault(storage StorageOrder, catalog product.CatalogProduct) (s *ServiceOrderDefault) {
	s = &ServiceOrderDefault{
		storage: storage,
		catalog: catalog,
	}
	return
}

// FindById returns the order with the given Id
func (s *ServiceOrderDefault) FindById(id int) (o Order, err error) {
	// search the order in the storage
	or, err := s.storage.FindById(id)
	if err != nil {
		err = ErrServiceOrderNotFound
		return
	}

	// serialize the order
	o.Products = make(map[product.Product]int, 0)
	for pr, qt := range or.Products {
		var p product.Product
		p, err = s.catalog.FindProductByName(pr)
		if err != nil {
			err = ErrServiceOrderProductNotFound
			return
		}

		o.Products[p] = qt
	}

	return
}

// Add adds an order to the storage
func (s *ServiceOrderDefault) Add(o *Order) (err error) {
	// deserialize the order
	or := OrderDB{
		Products: make(map[string]int, 0),
	}
	for pr, qt := range o.Products {
		or.Products[pr.Name()] = qt
	}

	// add the order to the storage
	err = s.storage.Add(&or)
	if err != nil {
		return
	}

	// update the order id
	o.Id = or.Id

	return
}

// Update updates the order with the given Id
func (s *ServiceOrderDefault) Update(o *Order) (err error) {
	// deserialize the order
	or := OrderDB{
		Id:       o.Id,
		Products: make(map[string]int, 0),
	}
	for pr, qt := range o.Products {
		or.Products[pr.Name()] = qt
	}

	// update the order in the storage
	err = s.storage.Update(&or)
	if err != nil {
		return
	}

	return
}