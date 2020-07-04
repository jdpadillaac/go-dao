package mysqld

import (
	"database/sql"
	"fmt"
	"log"

	"../../utilities"
	_ "github.com/go-sql-driver/mysql"
)

func get() *sql.DB {

	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	// user:Password@tcp(server:port)/database?tls=false&autocommit=true
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true", config.User, config.Password, config.Server, config.Port, config.Database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
	// dsn := fmt
}
