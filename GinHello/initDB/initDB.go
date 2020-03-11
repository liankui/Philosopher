package initDB

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var Db *gorm.DB

func init() {
	var err error
	//Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ginhello")
	Db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ginhello")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	//Db.SetMaxOpenConns(10)
	//Db.SetMaxIdleConns(10)
}
