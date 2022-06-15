package internal

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

const (
	dbUserName = "root"
	dbPassword = "Cricket@1626"
	dbHostName = "127.0.0.1:3306"
	dbName     = "cardb"
)

//Globals
var mysqlDb *sql.DB

func InitMysqlConnection() {

	var err error

	//username:password@tcp(hostIp:port)/dbName
	dbString := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUserName, dbPassword, dbHostName, dbName)

	mysqlDb, err = sql.Open("mysql", dbString)
	if err != nil {
		fmt.Println("Error connecting to DB: ", err)
		log.Error("Error connecting to DB: ", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to DB")

	//err = createDatabaseTables()
	//if err != nil {
	//	fmt.Println("Error creating table: ", err)
	//	log.Error("Error creating table: ", err)
	//	os.Exit(1)
	//}
}

func createDatabaseTables() (err error) {
	fmt.Println("Creating tables...")
	//create cars table
	_, err = mysqlDb.Exec(CARS_CREATE_TABLE_QUERY)
	if err != nil {
		fmt.Println("ERROR: create department table failed. ", err)
		log.Error("ERROR: create department table failed", err)
		return
	}

	fmt.Println("table created successfully")
	log.Info("table created successfully")
	return
}
