package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

func GetConnector() *sql.DB {
	cfg := mysql.Config{
		User:                 "유저 이름 입력",
		Passwd:               "비밀번호 입력",
		Net:                  "tcp",
		Addr:                 "<ip>:<포트>",
		Collation:            "utf8mb4_general_ci",
		Loc:                  time.UTC,
		MaxAllowedPacket:     4 << 20.,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		DBName:               "DB명 입력",
	}
	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		panic(err)
	}
	db := sql.OpenDB(connector)
	return db
}
