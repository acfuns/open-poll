package polls

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PollsRegister(router *gin.RouterGroup) {
	router.POST("/poll", CreateSurvey)
}

func CreateSurvey(c *gin.Context) {
	surveyModelValidator := SurveyModelValidator{}

	if err := surveyModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := SaveOne(&surveyModelValidator.surveyModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create survey",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Survey created successfully",
	})
}
