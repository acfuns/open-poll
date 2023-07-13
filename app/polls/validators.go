package polls

import "github.com/gin-gonic/gin"

type SurveyModelValidator struct {
	Title       string                   `json:"title" binding:"required"`
	Description string                   `json:"description"`
	Anonymous   bool                     `json:"anonymous"`
	Questions   []QuestionModelValidator `json:"questions" binding:"required"`

	surveyModel Survey `json:"-"`
}

type QuestionModelValidator struct {
	Text    string                 `json:"text" binding:"required"`
	Type    string                 `json:"type" binding:"required"`
	Options []OptionModelValidator `json:"options" binding:"required"`
}

type OptionModelValidator struct {
	Text string `json:"text" binding:"required"`
}

func (v *SurveyModelValidator) Bind(c *gin.Context) error {
	err := c.ShouldBindJSON(v)
	if err != nil {
		return err
	}
	v.surveyModel.UserID = c.GetUint("my_user_id")
	v.surveyModel.Title = v.Title
	v.surveyModel.Description = v.Description
	v.surveyModel.Anonymous = v.Anonymous
	for _, q := range v.Questions {
		question := Question{
			Text: q.Text,
			Type: q.Type,
		}

		for _, o := range q.Options {
			Option := Option{
				Text: o.Text,
			}
			question.Options = append(question.Options, Option)
		}
		v.surveyModel.Questions = append(v.surveyModel.Questions, question)
	}
	return nil
}
