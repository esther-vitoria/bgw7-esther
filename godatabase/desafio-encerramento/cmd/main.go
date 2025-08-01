package main

import (
	"app/internal/application"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Cria configuração MySQL a partir das envs
func buildMySQLConfig() *mysql.Config {
	return &mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  os.Getenv("DB_NET"),
		Addr:                 os.Getenv("DB_ADDR"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}
}

func main() {
	// Carregar variáveis do .env
	if err := godotenv.Load(); err != nil {
		log.Printf("Aviso: Não foi possível carregar .env (%v)", err)
	}

	mysqlCfg := buildMySQLConfig()

	// Configuração e execução da migração
	cfgMigrate := &application.ConfigApplicationMigrate{
		Db:               mysqlCfg,
		FilePathCustomer: os.Getenv("CUSTOMER_JSON"),
		FilePathProduct:  os.Getenv("PRODUCT_JSON"),
		FilePathInvoice:  os.Getenv("INVOICE_JSON"),
		FilePathSale:     os.Getenv("SALE_JSON"),
	}
	appMigrate := application.NewApplicationMigrate(cfgMigrate)
	defer appMigrate.TearDown()

	if err := appMigrate.SetUp(); err != nil {
		log.Fatalf("Erro na migração (SetUp): %v", err)
	}
	if err := appMigrate.Run(); err != nil {
		log.Fatalf("Erro na migração (Run): %v", err)
	}

	// Configuração e execução da API
	serverAddr := fmt.Sprintf("%s:%s",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
	)
	cfgDefault := &application.ConfigApplicationDefault{
		Db:   mysqlCfg,
		Addr: serverAddr,
	}
	app := application.NewApplicationDefault(cfgDefault)

	log.Printf("Servidor pronto em %s", serverAddr)
	if err := app.SetUp(); err != nil {
		log.Fatalf("Erro na API (SetUp): %v", err)
	}
	defer app.TearDown()

	if err := app.Run(); err != nil {
		log.Fatalf("Erro na API (Run): %v", err)
	}
}
