package catalog

import "bootcamp-web/internal"

// NewCatalogProductMap creates a new catalog of products.
func NewCatalogProductMap(db map[string]internal.Product) *CatalogProductMap {
	return &CatalogProductMap{db: db}
}

// CatalogProductMap is a struct that represents a catalog of products
type CatalogProductMap struct {
	// db is a map that stores a catalog of products (key: name, value: product)
	db map[string]internal.Product
}

// FindProductByName finds a product by its name.
func (s *CatalogProductMap) FindProductByName(name string) (p internal.Product, err error) {
	product, ok := s.db[name]
	if !ok {
		err = ErrCatalogProductNotFound
		return
	}

	p = product
	return
}