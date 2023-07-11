package users

import (
	"fmt"
	"net/http"
	"soul/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
)

// Strips `BEARER ` prefix from token string
func stripBearerPrefixFromTokenString(token string) (string, error) {
	fmt.Println(token)
	if len(token) > 5 && strings.ToUpper(token[0:7]) == "BEARER " {
		return token[7:], nil
	}
	return token, nil
}

// two token extractor

var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromTokenString,
}

var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

// A helper to write user_id and user_model to the context
func UpdateContextUserModel(c *gin.Context, my_user_id uint) {
	var myUserModel User
	if my_user_id != 0 {
		db := utils.GetDB()
		db.First(&myUserModel, my_user_id)
	}
	c.Set("my_user_id", my_user_id)
	c.Set("my_user_model", myUserModel)
}

func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// init context userid
		UpdateContextUserModel(c, 0)

		// access control
		if c.Request.URL.Path == "/api/login" {
			c.Next()
			return
		}

		if c.Request.URL.Path == "/api/register" {
			c.Next()
			return
		}

		token, err := request.ParseFromRequest(c.Request, MyAuth2Extractor, func(_ *jwt.Token) (interface{}, error) {
			b := ([]byte(utils.SecretPassword))
			return b, nil
		})
		if err != nil {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			my_user_id := uint(claims["id"].(float64))
			UpdateContextUserModel(c, my_user_id)
		}
	}
}
