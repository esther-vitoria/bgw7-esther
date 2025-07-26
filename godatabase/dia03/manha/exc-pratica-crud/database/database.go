package database

import (
	"database/sql"
	"fmt"

	"github.com/bgw7/exc-pratica-crud/config"
	_ "github.com/go-sql-driver/mysql"
)

var conn *sql.DB

func Connect() (err error) {
	cfg := config.NewConfig(nil)
	fmt.Println(*cfg.Server)

	connectionsParams := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		cfg.Server.Database.User,
		cfg.Server.Database.Password,
		cfg.Server.Database.Host,
		cfg.Server.Database.Port,
		cfg.Server.Database.Database,
	)

	fmt.Println(connectionsParams)

	db, err := sql.Open("mysql", connectionsParams)

	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}
	conn = db
	return
}

func GetDB() *sql.DB {
	return conn
}
