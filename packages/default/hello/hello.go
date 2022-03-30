package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xo/dburl"
)

func Main(args map[string]interface{}) map[string]interface{} {
	name, ok := args["name"].(string)
	if !ok {
		name = "stranger"
	}
	msg := make(map[string]interface{})
	msg["body"] = "Hello " + name + "!"

	// Open up our database connection.
	fmt.Println("Parsing db url....")
	connection := os.Getenv("DB_URL")
	dbURL, err := dburl.Parse(connection)
	if err != nil {
		panic(err)
	}

	dbPassword, _ := dbURL.User.Password()
	dbName := strings.Trim(dbURL.Path, "/")
	connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=true", dbURL.User.Username(), dbPassword, dbURL.Hostname(), dbURL.Port(), dbName)

	fmt.Println("Connecting to db....")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Pinging the db....")
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	msg["body"] = fmt.Sprintf("%s. MySQL Ping Successful", msg["body"])

	return msg
}
