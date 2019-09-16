package db

import (
	"github.com/rabbitmeow/golang-simple-api/config"
	// import dialect
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var conf = config.ReadConfig()

// Init is
func Init() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", conf.Database.User+":"+conf.Database.Password+"@tcp("+conf.Database.Host+":"+conf.Database.Port+")/"+conf.Database.Dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
}

// GetDB is
func GetDB() *gorm.DB {
	return db
}
