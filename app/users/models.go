package users

import (
	"errors"
	"soul/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"unique"`
	Password    string
	Email       string `gorm:"unique"`
	LastLogin   time.Time
	IsSuperuser bool
	IsActive    bool
}

func AutoMigrate() {
	db := utils.GetDB()

	db.AutoMigrate(&User{})
}

func (u *User) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}

func (u *User) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func SaveOne(data interface{}) error {
	db := utils.GetDB()
	err := db.Save(data).Error
	return err
}

func FindOneUser(condition interface{}) (User, error) {
	db := utils.GetDB()
	var model User
	err := db.Where(condition).First(&model).Error
	return model, err
}
