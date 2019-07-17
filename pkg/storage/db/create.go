package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type dbConfig struct {
	serverName string
	user       string
	password   string
	dbName     string
}

func CreateDatabase() (*sql.DB, error) {
	// Open json file containing configuration to open db
	file, err := os.Open("coupon_api/configs/config.json")
	if err != nil {
		log.Fatalln("Error opening config file:", err)
	}
	defer file.Close()

	// Decode json file and save results in struct
	decoder := json.NewDecoder(file)
	conf := dbConfig{}
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatalln("Error decoding config file:", err)
	}

	// Open db with config settings
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", conf.user, conf.password, conf.serverName, conf.dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
