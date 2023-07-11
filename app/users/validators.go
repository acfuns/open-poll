package users

import (
	"soul/utils"

	"github.com/gin-gonic/gin"
)

type UserModelValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"required,alpha,min=3,max=20"`
		Password string `form:"password" json:"password" binding:"required,min=8,max=20"`
		Email    string `form:"email" json:"email" binging:"required,email"`
	} `json:"user"`

	userModel User `json:"-"`
}

func (v *UserModelValidator) Bind(c *gin.Context) error {
	err := c.ShouldBind(v)
	if err != nil {
		return err
	}
	v.userModel.Username = v.User.Username
	v.userModel.Email = v.User.Email
	v.userModel.setPassword(v.User.Password)
	return nil
}

func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}
	return userModelValidator
}

type LoginValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"required,alpha,min=3,max=20"`
		Password string `form:"password" json:"password" binding:"required,min=8,max=20"`
	} `json:"user"`
	userModel User `json:"-"`
}

func (v *LoginValidator) Bind(c *gin.Context) error {
	err := utils.Bind(c, v)
	if err != nil {
		return err
	}

	v.userModel.Username = v.User.Username
	return nil
}

func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}
