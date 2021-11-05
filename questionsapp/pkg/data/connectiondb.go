package data

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	mgo "gopkg.in/mgo.v2"
)

// for mongodb
var db *mgo.Database

func GetMongoDB() *mgo.Database {
	host := "MONGO_HOST"
	dbName := "Go_mic_service"
	fmt.Println("connection info:", host, dbName)
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
	fmt.Println("session err:", err)
	os.Exit(2)
}
	db = session.DB(dbName)
	return db
}


// for mysql --- OPTIONAL --- 
var sqldb *sql.DB
func GetSqlDB() *sql.DB {
	var err error
	dbUser := "root" // mysql username
	dbPass := "root"  // mysql password
	dbName := "godb" // mysql dbname
	sqldb, err := sql.Open("mysql", dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
	panic(err.Error())
	}
	return sqldb
}


