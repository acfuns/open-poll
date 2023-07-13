package polls

import (
	"soul/utils"

	"gorm.io/gorm"
)

type Survey struct {
	gorm.Model
	UserID      uint   `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string
	Anonymous   bool `gorm:"default:false"`
	Questions   []Question
}

type Question struct {
	gorm.Model
	SurveyID uint   `gorm:"not null"`
	Text     string `gorm:"not null"`
	Type     string `gorm:"not null"`
	Options  []Option
}

type Option struct {
	gorm.Model
	QuestionID uint   `gorm:"not null"`
	Text       string `gorm:"not null"`
}

type Response struct {
	gorm.Model
	UserID     uint `gorm:"not null"`
	QuestionID uint `gorm:"not null"`
	OptionID   uint `gorm:"not null"`
}

func AutoMigrate() {
	db := utils.GetDB()
	db.AutoMigrate(&Survey{}, &Question{}, &Option{}, &Response{})
}

func SaveOne(data any) error {
	db := utils.GetDB()
	err := db.Save(data).Error
	return err
}
