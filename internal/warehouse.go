package internal

type WarehouseDB struct {
	// Id is the unique identifier of the warehouse
	Id int

	// Name is the name of the warehouse
	Name string

	// Stock is the stock of products in the warehouse with the quantity of each product
	// - key: product name
	// - value: quantity
	Stock map[string]int
}