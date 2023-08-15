package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": ctx.Request.Header,
	})
}
