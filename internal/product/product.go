package product

// Product is a struct that represents a product
type Product struct {
	name		string
	description	string
}

// Name returns the name of the product
func (p Product) Name() (name string) {
	name = p.name
	return
}

// Description returns the description of the product
func (p Product) Description() (description string) {
	description = p.description
	return
}