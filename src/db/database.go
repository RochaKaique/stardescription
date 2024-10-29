package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Driver
	"github.com/spf13/viper"
)

func Connect() (*sql.DB, error) {
	db, error := sql.Open("mysql", getConnectionString())
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		db.Close()
		return nil, error
	}

	return db, nil
}

func getConnectionString() string {
	return fmt.Sprintf("%s:%s@/%s?charset=utf8&&parseTime=True&loc=Local",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.schema"))
}
