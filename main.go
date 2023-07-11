package main

import (
	"soul/app/users"
	"soul/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Migrate() {
	users.AutoMigrate()
}

func main() {
	r := gin.Default()

	// // config log
	// logger, _ := zap.NewProduction()
	// defer logger.Sync()
	// sugar := logger.Sugar()
	utils.Init()

	Migrate()

	r.Use(cors.Default())
	v1 := r.Group("/api")
	v1.Use(users.AuthMiddleware(true))
	users.UsersRegister(v1)

	r.Run("127.0.0.1:8000")
}
