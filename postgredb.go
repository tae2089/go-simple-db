package main

import (
	"database/sql"
	"fmt"
	"log"
)

func GetPostGresSqlConnetion(dbconfig *DBConfig) *sql.DB {
	query := fmt.Sprintf("host=%s port=%d  user=%s password=%s dbname=%s sslmode=disable",
		dbconfig.Host, dbconfig.Port, dbconfig.User, dbconfig.Password, dbconfig.DbName)
	//db, err := postsql.Open(query)
	db, err := sql.Open("postgresql", query)
	if err != nil || db.Ping() != nil {
		panic(err)
	}
	return db
}

//func GetDataByQuery(db *sql.DB, datas map[string]interface{}) {
//	for key, value := range datas {
//		log.Println(key, value)
//	}
//}

func GetDataByQuery(data map[string]interface{}) {
	for key, value := range data {
		log.Println(key, value)
	}
}
func GetDataByQueryRow() {

}
