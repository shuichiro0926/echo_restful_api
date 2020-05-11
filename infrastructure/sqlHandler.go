package infrastructure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func DBInit() {
	db, err = gorm.Open("mysql",  "root:root@tcp(127.0.0.1:3308)/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}
}

func GetConnection() *gorm.DB {
	return db
}
