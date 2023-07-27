package polls

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PollsRegister(router *gin.RouterGroup) {
	router.POST("/poll", CreateSurvey)
	router.GET("/poll", ListSurvey)
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

func ListSurvey(c *gin.Context) {
	// todo: query survey with the user from database
	userId := c.GetUint("user_id")
	survey, err := FindAllSurveyWithUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to find survey",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"survey": survey,
	})
}
