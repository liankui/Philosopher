package initDB

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ginhello")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}
