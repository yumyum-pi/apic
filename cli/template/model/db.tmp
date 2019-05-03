package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB is a pointer to database
var DB *sql.DB

const dbDriver = "mysql"

var dbUser, dbPass, dbName string

// Connect function connects to database and return db interface
func connect() {
	connectionString := fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName)
	var err error

	DB, err = sql.Open(dbDriver, connectionString)
	if err != nil {
		panic(err.Error())
	}
}

// Init Initialization of databse
// get username, userpassword and database name
func Init() {
	// load environment variables from .evn file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setting the variable from environment
	dbUser = os.Getenv("MYSQL_USER")
	dbPass = os.Getenv("MYSQL_PASSWORD")
	dbName = os.Getenv("MYSQL_DBNAME")

	connect()
}

// GetDB returns a pointer of Database
func GetDB() *sql.DB {
	return DB
}
