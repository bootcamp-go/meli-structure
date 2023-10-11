package handler

import (
	"bootcamp-web/internal"
	"bootcamp-web/internal/storage"
	"bootcamp-web/platform/web"
	"bootcamp-web/platform/web/middlewares"
	"net/http"
)

// NewHandlerWarehouse creates a new handler of a warehouse
func NewHandlerWarehouse(st storage.StorageWarehouse) *HandlerWarehouse {
	return &HandlerWarehouse{st: st}
}

// HandlerWarehouse is an struct that represents the handler of a warehouse
type HandlerWarehouse struct {
	// storage is the storage of the warehouse
	st storage.StorageWarehouse
}

// AddWarehouse is a handler that adds a warehouse to the storage
type RequestAddWarehouse struct {
	// name is the name of the warehouse
	Name string			 `json:"name"`

	// Stock is the stock of products in the warehouse with the quantity of each product
	// - key: product name
	// - value: quantity of the product
	Stock map[string]int `json:"stock"`
}
func (h *HandlerWarehouse) AddWarehouse() web.Handler {
	return func(w http.ResponseWriter, r *http.Request) (err error) {
		// request
		// - body: validate required client fields dynamically
		err = web.ValidatorRequiredJSON(r.Body, "name", "stock")
		if err != nil {
			err = middlewares.NewError(http.StatusBadRequest, "missing required fields")
			return
		}
		// - body: decode request body
		var req RequestAddWarehouse
		err = web.DecodeJSON(r, &req)
		if err != nil {
			err = middlewares.NewError(http.StatusBadRequest, "invalid request body")
			return
		}

		// process
		// - create warehouse
		wh := internal.WarehouseDB{
			Name: req.Name,
			Stock: req.Stock,
		}
		err = h.st.Add(&wh)
		if err != nil {
			err = middlewares.NewError(http.StatusInternalServerError, "error adding warehouse")
			return
		}

		// response
		web.EncodeJSON(w, map[string]any{"message": "warehouse added successfully", "data": wh.Id}, http.StatusCreated)
		return
	}
}

// AddProductStock is a handler that adds a product to the stock of a warehouse
type RequestAddProductStock struct {
	// name is the name of the product
	Name string `json:"name"`
	// quantity is the quantity of the product
	Quantity int `json:"quantity"`
}
func (h *HandlerWarehouse) AddProductStock() web.Handler {
	return func(w http.ResponseWriter, r *http.Request) (err error) {
		// request
		// - body: validate required client fields dynamically
		err = web.ValidatorRequiredJSON(r.Body, "name", "quantity")
		if err != nil {
			err = middlewares.NewError(http.StatusBadRequest, "missing required fields")
			return
		}
		// - body: decode request body
		var req RequestAddProductStock
		err = web.DecodeJSON(r, &req)
		if err != nil {
			err = middlewares.NewError(http.StatusBadRequest, "invalid request body")
			return
		}
		// - path-param: get warehouse id
		whId, err := web.ParamInt(r, "id")
		if err != nil {
			err = middlewares.NewError(http.StatusBadRequest, "invalid warehouse id")
			return
		}

		// process
		// - get warehouse
		wh, err := h.st.FindById(whId)
		if err != nil {
			err = middlewares.NewError(http.StatusInternalServerError, "error getting warehouse")
			return
		}
		// - add product to stock
		wh.Stock[req.Name] += req.Quantity
		// - update warehouse
		err = h.st.Update(&wh)
		if err != nil {
			err = middlewares.NewError(http.StatusInternalServerError, "error updating warehouse")
			return
		}

		// response
		web.EncodeJSON(w, map[string]any{"message": "product added to warehouse successfully"}, http.StatusOK)
		return
	}
}