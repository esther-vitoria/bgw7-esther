package application

import (
	"app/internal/customer"
	"app/internal/domain"
	"app/internal/invoice"
	"app/internal/loader"
	"app/internal/migrator"
	"app/internal/product"
	"app/internal/sale"
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

// ConfigApplicationMigrate is the struct that contains the paths to the files that will be loaded
type ConfigApplicationMigrate struct {
	Db               *mysql.Config
	FilePathCustomer string
	FilePathProduct  string
	FilePathInvoice  string
	FilePathSale     string
}

// NewApplicationMigrate returns a new ApplicationMigrate
func NewApplicationMigrate(config *ConfigApplicationMigrate) (a *ApplicationMigrate) {
	a = &ApplicationMigrate{
		config: config,
	}
	return
}

// ApplicationMigrate is the implementation of the interface ApplicationMigrate
type ApplicationMigrate struct {
	// config is the configuration of the application
	config *ConfigApplicationMigrate
	// database is the database to load the data
	database *sql.DB
	// fileCustomer is the path to the file that contains the customers
	fileCustomer *os.File
	// fileProduct is the path to the file that contains the products
	fileProduct *os.File
	// fileInvoice is the path to the file that contains the invoices
	fileInvoice *os.File
	// fileSales is the path to the file that contains the sales
	fileSales *os.File
	// Migrators
	migrators []domain.Migrator
}

// TearDown is the method to tear down the application migrate
func (a *ApplicationMigrate) TearDown() {
	// - close files
	if a.fileCustomer != nil {
		a.fileCustomer.Close()
	}
	if a.fileProduct != nil {
		a.fileProduct.Close()
	}
	if a.fileInvoice != nil {
		a.fileInvoice.Close()
	}
	if a.fileSales != nil {
		a.fileSales.Close()
	}
	// - close db
	if a.database != nil {
		a.database.Close()
	}
}

// SetUp is the method to set up the application migrate
func (a *ApplicationMigrate) SetUp() (err error) {
	// dependencies
	// - db: init
	a.database, err = sql.Open("mysql", a.config.Db.FormatDSN())
	if err != nil {
		return
	}
	// - db: ping
	err = a.database.Ping()
	if err != nil {
		return
	}
	// - file
	a.fileCustomer, err = os.Open(a.config.FilePathCustomer)
	if err != nil {
		return
	}
	a.fileProduct, err = os.Open(a.config.FilePathProduct)
	if err != nil {
		return
	}
	a.fileInvoice, err = os.Open(a.config.FilePathInvoice)
	if err != nil {
		return
	}
	a.fileSales, err = os.Open(a.config.FilePathSale)
	if err != nil {
		return
	}
	// - migrators
	ldCustomer := loader.NewCustomersJSON(a.fileCustomer)
	rpCustomer := customer.NewCustomersMySQL(a.database)
	mgCustomer := migrator.NewMigratorCustomerToDatabase(ldCustomer, rpCustomer)

	ldProduct := loader.NewProductsJSON(a.fileProduct)
	rpProduct := product.NewProductsMySQL(a.database)
	mgProduct := migrator.NewMigratorProductToDatabase(ldProduct, rpProduct)

	ldInvoice := loader.NewInvoicesJSON(a.fileInvoice)
	rpInvoice := invoice.NewInvoicesMySQL(a.database)
	mgInvoice := migrator.NewMigratorInvoiceToDatabase(ldInvoice, rpInvoice)

	ldSale := loader.NewSalesJSON(a.fileSales)
	rpSale := sale.NewSalesMySQL(a.database)
	mgSale := migrator.NewMigratorSaleToDatabase(ldSale, rpSale)

	a.migrators = []domain.Migrator{
		mgCustomer,
		mgInvoice,
		mgProduct,
		mgSale,
	}

	return
}

// Run is the method to run the application migrate
func (a *ApplicationMigrate) Run() (err error) {
	for _, v := range a.migrators {
		err = v.Migrate()
		if err != nil {
			return
		}
	}

	return
}
