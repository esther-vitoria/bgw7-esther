package application

import (
	"app/handler"
	"app/internal/customer"
	"app/internal/invoice"
	"app/internal/product"
	"app/internal/sale"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
)

// ConfigApplicationDefault is the configuration for NewApplicationDefault.
type ConfigApplicationDefault struct {
	// Db is the database configuration.
	Db *mysql.Config
	// Addr is the server address.
	Addr string
}

// NewApplicationDefault creates a new ApplicationDefault.
func NewApplicationDefault(config *ConfigApplicationDefault) *ApplicationDefault {
	// default values
	defaultCfg := &ConfigApplicationDefault{
		Db:   nil,
		Addr: ":8080",
	}
	if config != nil {
		if config.Db != nil {
			defaultCfg.Db = config.Db
		}
		if config.Addr != "" {
			defaultCfg.Addr = config.Addr
		}
	}

	return &ApplicationDefault{
		cfgDb:   defaultCfg.Db,
		cfgAddr: defaultCfg.Addr,
	}
}

// ApplicationDefault is an implementation of the Application interface.
type ApplicationDefault struct {
	// cfgDb is the database configuration.
	cfgDb *mysql.Config
	// cfgAddr is the server address.
	cfgAddr string
	// db is the database connection.
	db *sql.DB
	// router is the chi router.
	router *chi.Mux
}

// TearDown tears down the application.
func (a *ApplicationDefault) TearDown() {
	// close db
	if a.db != nil {
		a.db.Close()
	}
}

// SetUp sets up the application.
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	// - db: init
	a.db, err = sql.Open("mysql", a.cfgDb.FormatDSN())
	if err != nil {
		return
	}
	// - db: ping
	err = a.db.Ping()
	if err != nil {
		return
	}
	// - repository
	rpCustomer := customer.NewCustomersMySQL(a.db)
	rpProduct := product.NewProductsMySQL(a.db)
	rpInvoice := invoice.NewInvoicesMySQL(a.db)
	rpSale := sale.NewSalesMySQL(a.db)
	// - service
	svCustomer := customer.NewCustomersDefault(rpCustomer)
	svProduct := product.NewProductsDefault(rpProduct)
	svInvoice := invoice.NewInvoicesDefault(rpInvoice)
	svSale := sale.NewSalesDefault(rpSale)
	// - handler
	hdCustomer := handler.NewCustomersDefault(svCustomer)
	hdProduct := handler.NewProductsDefault(svProduct)
	hdInvoice := handler.NewInvoicesDefault(svInvoice)
	hdSale := handler.NewSalesDefault(svSale)

	// routes
	// - router
	a.router = chi.NewRouter()
	// - middlewares
	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Recoverer)
	// - endpoints
	a.router.Route("/customers", func(r chi.Router) {
		// - GET /customers
		r.Get("/", hdCustomer.GetAll())
		// - GET /customers/top-active
		r.Get("/top-active", hdCustomer.GetTopActiveCustomersByAmountSpent())
		// - GET /customers/invoices-by-condition
		r.Get("/invoices-by-condition", hdCustomer.GetInvoicesByCondition())
		// - POST /customers
		r.Post("/", hdCustomer.Create())
	})
	a.router.Route("/products", func(r chi.Router) {
		// - GET /products
		r.Get("/", hdProduct.GetAll())
		// - GET /products/top-sold
		r.Get("/top-sold", hdProduct.GetTopProductsByAmountSold())
		// - POST /products
		r.Post("/", hdProduct.Create())
	})
	a.router.Route("/invoices", func(r chi.Router) {
		// - GET /invoices
		r.Get("/", hdInvoice.GetAll())
		// - POST /invoices
		r.Post("/", hdInvoice.Create())
		// - PUT /invoices/total
		r.Put("/total", hdInvoice.UpdateAllTotal())
	})
	a.router.Route("/sales", func(r chi.Router) {
		// - GET /sales
		r.Get("/", hdSale.GetAll())
		// - POST /sales
		r.Post("/", hdSale.Create())
	})

	return
}

// Run runs the application.
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.cfgAddr, a.router)
	return
}
