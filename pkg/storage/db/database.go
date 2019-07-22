package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	ServerName string
	User       string
	Password   string
	DbName     string
}

func Open() (*sql.DB, error) {
	// Open json file containing configuration to open db
	file, err := os.Open("../../configs/config.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
	}
	defer file.Close()

	// Decode json file and save results in struct
	decoder := json.NewDecoder(file)
	conf := config{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error decoding config file:", err)
	}

	// Open mysql without specifying database as admin
	mysqlConnection := fmt.Sprintf("%s:%s@tcp(%s)/", conf.User, conf.Password, conf.ServerName)
	db, err := sql.Open("mysql", mysqlConnection)
	if err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
		return nil, err
	}

	// Create database if it doesn't exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + conf.DbName)
	if err != nil {
		panic(err)
	}
	db.Close()

	// Connect to selected database
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.User, conf.Password, conf.ServerName, conf.DbName)
	db, err = sql.Open("mysql", dbConnection)
	if err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
		return nil, err
	}

	return db, nil
}
