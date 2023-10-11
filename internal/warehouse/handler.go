package warehouse

import (
	"bootcamp-web/platform/web"
	"bootcamp-web/platform/web/validator"
	"net/http"
)

// HandlerWarehouse is an struct that represents the handler of a warehouse
type HandlerWarehouse struct {
	// storage is the storage of the warehouse
	st StorageWarehouse
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
func (h *HandlerWarehouse) AddWarehouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body: validate required client fields dynamically
		err := validator.RequiredJSON(r.Body, "name", "stock")
		if err != nil {
			web.EncodeJSON(w, map[string]any{"message": "missing required fields"}, http.StatusBadRequest)
			return
		}
		// - body: decode request body
		var req RequestAddWarehouse
		err = web.DecodeJSON(r, &req)
		if err != nil {
			web.EncodeJSON(w, map[string]any{"message": "invalid request body"}, http.StatusBadRequest)
			return
		}

		// process
		// - create warehouse
		wh := WarehouseDB{
			Name: req.Name,
			Stock: req.Stock,
		}
		err = h.st.Add(&wh)
		if err != nil {
			web.EncodeJSON(w, map[string]any{"message": "error adding warehouse"}, http.StatusInternalServerError)
			return
		}

		// response
		web.EncodeJSON(w, map[string]any{"message": "warehouse added successfully", "data": wh.Id}, http.StatusCreated)
	}
}

// AddProductStock is a handler that adds a product to the stock of a warehouse
type RequestAddProductStock struct {
	// name is the name of the product
	Name string `json:"name"`
	// quantity is the quantity of the product
	Quantity int `json:"quantity"`
}
func (h *HandlerWarehouse) AddProductStock() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body: validate required client fields dynamically
		err := validator.RequiredJSON(r.Body, "name", "quantity")
		if err != nil {
			web.EncodeJSON(w, map[string]any{"message": "missing required fields"}, http.StatusBadRequest)
			return
		}
		// - body: decode request body
		var req RequestAddProductStock
		err = web.DecodeJSON(r, &req)
		if err != nil {
			web.EncodeJSON(w, map[string]any{"message": "invalid request body"}, http.StatusBadRequest)
			return
		}
		// - path-param: get warehouse id
		whId, err := web.ParamInt(r, "id")
		if err != nil {
			web.EncodeJSON(w, map[string]any{"message": "invalid warehouse id"}, http.StatusBadRequest)
			return
		}

		// process
		// - get warehouse
		wh, err := h.st.FindById(whId)
		if err != nil {
			web.EncodeJSON(w, map[string]any{"message": "error getting warehouse"}, http.StatusInternalServerError)
			return
		}
		// - add product to stock
		wh.Stock[req.Name] += req.Quantity
		// - update warehouse
		err = h.st.Update(&wh)
		if err != nil {
			web.EncodeJSON(w, map[string]any{"message": "error updating warehouse"}, http.StatusInternalServerError)
			return
		}

		// response
		web.EncodeJSON(w, map[string]any{"message": "product added to warehouse successfully"}, http.StatusOK)
	}
}