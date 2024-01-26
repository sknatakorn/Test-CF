package infrastructure

import (
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {
	Db = connectToSQLite()
}

func connectToSQLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(viper.GetString("db.name")), &gorm.Config{})
	if err != nil {
		panic("Cannot Connect to db")
	}

	return db
}
