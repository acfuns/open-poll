package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Database struct {
// 	*gorm.DB
// }

// register global DB
var DB *gorm.DB

func Init() *gorm.DB {
	// env var config
	cfg, err := Cfg()
	if err != nil {
		panic(err)
	}

	// config database
	dsn := fmt.Sprintf("host=%s user=%s password=%d dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", cfg.HOST, cfg.USER, cfg.PASSWORD, cfg.DB_NAME, cfg.PORT)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init)", err)
	}

	return DB
}

// util: get global db
func GetDB() *gorm.DB {
	return DB
}
