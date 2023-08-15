package utils

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// type Database struct {
// 	*gorm.DB
// }

// register global DB
var DB *gorm.DB

func Init() *gorm.DB {
	// env var config
	// cfg, err := Cfg()
	// if err != nil {
	// 	panic(err)
	// }

	// config database
	var err error
	DB, err = gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init)", err)
	}

	return DB
}

// util: get global db
func GetDB() *gorm.DB {
	return DB
}
