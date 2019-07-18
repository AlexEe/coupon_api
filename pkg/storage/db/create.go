package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type dbConfig struct {
	ServerName string
	User       string
	Password   string
	DbName     string
}

func CreateDatabase() (*sql.DB, error) {
	// Open json file containing configuration to open db
	// path, _ := filepath.Abs("../../configs/config.json")
	file, err := os.Open("../../configs/config.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
	}

	defer file.Close()

	// Decode json file and save results in struct
	byteValue, _ := ioutil.ReadAll(file)
	// decoder := json.NewDecoder(file)
	conf := dbConfig{}
	json.Unmarshal(byteValue, &conf)
	// err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error decoding config file:", err)
	}

	// Open db with config settings
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.User, conf.Password, conf.ServerName, conf.DbName)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
