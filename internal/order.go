package internal

// Order is a struct that represents an order
type Order struct {
	// Name is the name of the order
	Name string
	// Products is the list of products in the order with the quantity of each product
	// - key: product name
	// - value: quantity
	Products map[string]int
}