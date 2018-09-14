package config

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/configor"

	 _ "github.com/jinzhu/gorm/dialects/postgres"
)

type dbConfig struct{
	Host string
	Port string
	User string
	Password string
	DbName string
}

var _dbConn *gorm.DB

func MustGetDB() *gorm.DB {
	if _dbConn != nil {
		return _dbConn
	}

	var dbconf dbConfig
	err := configor.New(&configor.Config{ENVPrefix: "DB"}).Load(&dbconf)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open("postgres", "host="+dbconf.Host+" port="+dbconf.Port+" user="+dbconf.User+" dbname="+dbconf.DbName+" password="+dbconf.Password+" sslmode=disable")
	if err != nil {
		panic(err)
	}
	_dbConn=db
	return _dbConn
}