package product

// CatalogProductMap is a struct that represents a catalog of products
type CatalogProductMap struct {
	// db is a map that stores a catalog of products (key: name, value: product)
	db map[string]Product
}

// FindProductByName finds a product by its name.
func (s *CatalogProductMap) FindProductByName(name string) (p Product, err error) {
	product, ok := s.db[name]
	if !ok {
		err = ErrCatalogProductNotFound
		return
	}

	p = product
	return
}