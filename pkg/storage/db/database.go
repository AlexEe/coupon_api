package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	ServerName string
	User       string
	Password   string
	DbName     string
	Table      string
}

func Open() (*sql.DB, error) {
	// Open json file containing configuration to open db
	file, err := os.Open("../../configs/config.json")
	checkErr(err, "Error opening config file:")
	defer file.Close()

	// Decode json file and save results in struct
	decoder := json.NewDecoder(file)
	conf := config{}
	err = decoder.Decode(&conf)
	checkErr(err, "Error opening new database:")

	// Open mysql without specifying database as admin
	mysqlConnection := fmt.Sprintf("%s:%s@tcp(%s)/", conf.User, conf.Password, conf.ServerName)
	db, err := sql.Open("mysql", mysqlConnection)
	checkErr(err, "Error opening new database:")

	// Create database if it doesn't exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + conf.DbName)
	checkErr(err, "Error opening database:")
	db.Close()

	// Open new database
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.User, conf.Password, conf.ServerName, conf.DbName)
	db, err = sql.Open("mysql", dbConnection)
	checkErr(err, "Error opening new database:")

	// Use new database
	_, err = db.Exec("USE " + conf.DbName)
	checkErr(err, "Error using database:")

	// Create table if it doesn't exist
	// "coupons(id INT(10) NOT NULL AUTO_INCREMENT, name VARCHAR(150) NOT NULL, brand VARCHAR(150) NOT NULL, value DECIMAL(6,2) NOT NULL, created DATE NOT NULL, expiry DATE NOT NULL, PRIMARY KEY (id));"
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS " + conf.Table)
	checkErr(err, "Error creating table:")
	_, err = stmt.Exec()
	checkErr(err, "Error creating table:")

	return db, nil
}

func checkErr(err error, message string) {
	if err != nil {
		log.Fatalln(message, err)
	}
}
