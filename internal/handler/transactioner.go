package handler

import (
	"bootcamp-web/internal"
	"bootcamp-web/internal/transactioner"
	"bootcamp-web/platform/web"
	"bootcamp-web/platform/web/middlewares"
	"net/http"
)

// NewHandlerTransactioner creates a new handler of a transactioner
func NewHandlerTransactioner(tr transactioner.Transactioner) *HandlerTransactioner {
	return &HandlerTransactioner{tr: tr}
}

// HandlerTransactioner is an struct that represents the handler of a transactioner
type HandlerTransactioner struct {
	// tr is the transactioner
	tr transactioner.Transactioner
}

type ProductsJSON struct {
	// Name is the name of the product
	Name string `json:"name"`
	// Quantity is the quantity of the product
	Quantity int `json:"quantity"`
}
// RequestFulfill is a handler that fulfills a transaction between an order and a warehouse
type RequestFulfill struct {
	// Name is the name of the order
	Name string `json:"name"`

	// WarehouseName is the name of the warehouse
	WarehouseName string `json:"warehouse_name"`

	// Products is an slice of Product
	Products []ProductsJSON `json:"products"`
}
func (h *HandlerTransactioner) Fulfill() web.Handler {
	return func(w http.ResponseWriter, r *http.Request) (err error) {
		// request
		// - body: validate required client fields dynamically
		err = web.ValidatorRequiredJSON(r.Body, "name", "warehouse_name", "products")
		if err != nil {
			err = middlewares.NewError(http.StatusBadRequest, "missing required fields")
			return
		}
		// - body: decode request body
		var req RequestFulfill
		err = web.DecodeJSON(r, &req)
		if err != nil {
			err = middlewares.NewError(http.StatusBadRequest, "invalid request body")
			return
		}

		// process the transaction
		// - deserialize the order
		order := internal.Order{
			Name:     req.Name,
			Products: make(map[string]int),
		}
		for _, pr := range req.Products {
			order.Products[pr.Name] = pr.Quantity
		}
		// - fulfill the order
		err = h.tr.Fullfill(order, req.WarehouseName)
		if err != nil {
			switch err {
			case transactioner.ErrTransactionerOrderQuantityNotPositive:
				err = middlewares.NewError(http.StatusUnprocessableEntity, "quantity of a product in the order is not positive")
			case transactioner.ErrTransactionerOrderQuantityNotAvailable:
				err = middlewares.NewError(http.StatusUnprocessableEntity, "quantity of a product in the order is not available in the warehouse")
			case transactioner.ErrTransactionerWarehouseNotFound:
				err = middlewares.NewError(http.StatusUnprocessableEntity, "warehouse not found")
			case transactioner.ErrTransactionerWarehouseProductNotFound:
				err = middlewares.NewError(http.StatusUnprocessableEntity, "product not found in the warehouse")
			case transactioner.ErrTransactionerWarehouseProductQuantityInvalid:
				err = middlewares.NewError(http.StatusUnprocessableEntity, "invalid quantity of a product in the warehouse")
			default:
				err = middlewares.NewError(http.StatusInternalServerError, "internal server error")
			}
			return
		}
		
		// response
		// - serialize the order
		data := make([]ProductsJSON, 0, len(order.Products))
		for name, quantity := range order.Products {
			data = append(data, ProductsJSON{
				Name:     name,
				Quantity: quantity,
			})
		}
		web.EncodeJSON(w, map[string]any{"message": "order fulfilled", "data": data}, http.StatusOK)
		return
	}
}