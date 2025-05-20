package dbquery

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func ConnectToDb(user string, pass string, addr string, dbName string) error {
	cfg := mysql.Config{
		User:                 user,
		Passwd:               pass,
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               dbName,
		AllowNativePasswords: true,
		AllowOldPasswords:    true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return err
	}

	err = db.Ping()

	if err != nil {
		return err
	}

	return nil
}
