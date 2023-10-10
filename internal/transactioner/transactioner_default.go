package transactioner

import (
	"bootcamp-web/internal/order"
	"bootcamp-web/internal/warehouse"
)

type TransactionerDefault struct {
	// serviceWarehouse is the service to manage warehouses
	serviceWarehouse warehouse.ServiceWarehouse

	// serviceOrder is the service to manage orders
	serviceOrder order.ServiceOrder
}

// NewTransactionerDefault returns a new instance of TransactionerDefault
func NewTransactionerDefault(serviceWarehouse warehouse.ServiceWarehouse, serviceOrder order.ServiceOrder) (t *TransactionerDefault) {
	t = &TransactionerDefault{
		serviceWarehouse: serviceWarehouse,
		serviceOrder:     serviceOrder,
	}
	return
}

// Fulfill processes the transaction between the order and the warehouse
func (t *TransactionerDefault) Fullfill(orderId, warehouseId int) (err error) {
	// get the order
	or, err := t.serviceOrder.FindById(orderId)
	if err != nil {
		err = ErrTransactionerOrderNotFound
		return
	}

	// get the warehouse
	wh, err := t.serviceWarehouse.FindById(warehouseId)
	if err != nil {
		err = ErrTransactionerWarehouseNotFound
		return
	}

	// process the transaction
	for prOr, qt := range or.Products {
		prWh, ok := wh.Attributes.Stock[prOr]
		if !ok {
			err = ErrTransactionerWarehouseProductNotFound
			return
		}

		if prWh < qt {
			err = ErrTransactionerOrderQuantityNotAvailable
			return
		}
			
		wh.Attributes.Stock[prOr] = prWh - qt
	}

	// commit the transaction
	err = t.serviceWarehouse.Update(&wh)
	if err != nil {
		return
	}

	return
}