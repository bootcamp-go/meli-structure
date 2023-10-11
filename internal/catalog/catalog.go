package catalog

import (
	"bootcamp-web/internal"
	"errors"
)

var (
	// ErrCatalogProductNotFound is an error that occurs when a product is not found in the storage
	ErrCatalogProductNotFound = errors.New("catalog-product error: product not found")
)

// CatalogProduct is an interface that represents a catalog of products
type CatalogProduct interface {
	// FindProductByName finds a product by its name.
	FindProductByName(name string) (p internal.Product, err error)
}