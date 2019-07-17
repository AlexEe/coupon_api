package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func openDatabase() (*sql.DB, error) {
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/coupon")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	return db, err
}
