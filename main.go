package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//db := GetConnector()
	db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	var name string
	err = db.QueryRow("SELECT name FROM test WHERE id = 1").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
	print("success")
}
