package transactioner

import (
	"bootcamp-web/internal/warehouse"
	"errors"
)

type TransactionerDefault struct {
	// storageWarehouse is the storage of the warehouse
	storageWarehouse warehouse.StorageWarehouse
}

// NewTransactionerDefault returns a new instance of TransactionerDefault
func NewTransactionerDefault(storageWarehouse warehouse.StorageWarehouse) (t *TransactionerDefault) {
	t = &TransactionerDefault{
		storageWarehouse: storageWarehouse,
	}
	return
}

// Fulfill processes the transaction between the order and the warehouse
func (t *TransactionerDefault) Fullfill(order Order, warehouseName string) (err error) {
	// validation
	// - check if order quantity is positive
	for _, qt := range order.Products {
		if qt <= 0 {
			err = ErrTransactionerOrderQuantityNotPositive
			return
		}
	}
	
	// get the warehouse
	wh, err := t.storageWarehouse.FindByName(warehouseName)
	if err != nil {
		err = ErrTransactionerWarehouseNotFound
		return
	}

	// process the transaction
	for prOr, qtOr := range order.Products {
		// check if product is in the warehouse
		qtWh, ok := wh.Stock[prOr]
		if !ok {
			err = ErrTransactionerWarehouseProductNotFound
			return
		}

		// check if product quantity is available in the warehouse
		if qtWh < qtOr {
			err = ErrTransactionerOrderQuantityNotAvailable
			return
		}

		// update the product quantity in the warehouse
		qtWh -= qtOr
		wh.Stock[prOr] = qtWh
	}

	// update the warehouse
	err = t.storageWarehouse.Update(&wh)
	if err != nil {
		switch {
		case errors.Is(err, warehouse.ErrStorageWarehouseProductNotFound):
			err = ErrTransactionerWarehouseProductNotFound
		case errors.Is(err, warehouse.ErrStorageWarehouseInvalidQuantity):
			err = ErrTransactionerWarehouseProductQuantityInvalid
		}
		return
	}

	return
}