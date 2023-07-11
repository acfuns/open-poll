package test

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(ctx *gin.Context) {
	fmt.Println(ctx.Request.Header)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
