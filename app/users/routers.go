package users

import (
	"errors"
	"net/http"
	"soul/api/test"
	"soul/utils"

	"github.com/gin-gonic/gin"
)

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.GET("/ping", test.Pong)
}

func Register(c *gin.Context) {
	userModelValidator := NewUserModelValidator()

	// parse request body
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := SaveOne(&userModelValidator.userModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}

func Login(c *gin.Context) {
	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.NewValidatorError(err))
		return
	}

	userModel, err := FindOneUser(&User{Username: loginValidator.userModel.Username})
	if err != nil {
		c.JSON(http.StatusForbidden, utils.NewError("login", errors.New("not Registered username or invalid password")))
		return
	}

	if userModel.checkPassword(loginValidator.User.Password) != nil {
		c.JSON(http.StatusForbidden, utils.NewError("login", errors.New("not Registered username or invalid password")))
		return
	}
	UpdateContextUserModel(c, userModel.ID)
	c.JSON(http.StatusOK, gin.H{
		"token": utils.GenToken(userModel.ID),
	})
}
