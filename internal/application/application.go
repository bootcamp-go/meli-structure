package application

import (
	"bootcamp-web/internal/product"
	"bootcamp-web/internal/transactioner"
	"bootcamp-web/internal/warehouse"
	"bootcamp-web/platform/web"
	"bootcamp-web/platform/web/middlewares"
	"context"
	"log"
	"net"
	"net/http"
	"os"
)

// Application represents the application running in a http server
// with all its dependencies.
type Application struct {
	server  *http.Server
	network string
	address string
}

// New creates a new un-started application.
func New() *Application {
	muxer := web.NewMux(middlewares.NewErrorMiddleware(), middlewares.NewPanic())
	registerRoutes(muxer)
	httpServer := &http.Server{Handler: muxer}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Application{
		server:  httpServer,
		network: "tcp",
		address: ":" + port,
	}
}

// Start starts the application and blocks until the application is stopped.
func (a *Application) Start() error {
	ln, err := net.Listen(a.network, a.address)
	if err != nil {
		return err
	}

	log.Printf("Started at %s\n", ln.Addr().String())

	return a.server.Serve(ln)
}

// Stop stops the application.
func (a *Application) Stop() error {
	return a.server.Shutdown(context.Background())
}

func registerRoutes(m *web.Muxer) {
	// dependencies
	// - products
	catalogProducts := product.NewCatalogProductMap(
		make(map[string]product.Product),
	)
	// - warehouse
	warehouseStorage := warehouse.NewStorageWarehouseDefaultValidator(
		warehouse.NewStorageWarehouseCatalogValidator(
			warehouse.NewStorageWarehouseMap(
				make(map[int]warehouse.WarehouseDB),
				0,
			),
			catalogProducts,
		),
	)
	warehouseHandler := warehouse.NewHandlerWarehouse(warehouseStorage)

	// - transactioner
	tr := transactioner.NewTransactionerDefault(warehouseStorage)
	trHandler := transactioner.NewHandlerTransactioner(tr)

	// routes
	// - warehouse
	m.Handle("POST", "/warehouses", warehouseHandler.AddWarehouse())
	m.Handle("POST", "/warehouses/{warehouse_id}/products", warehouseHandler.AddProductStock())
	// - transactioner
	m.Handle("POST", "/orders", trHandler.Fulfill())
}