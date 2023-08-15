package users

import (
	"net/http"
	"soul/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// A helper to write user_id and user_model to the context
func UpdateContextUserModel(c *gin.Context, user_id uint) {
	var myUserModel User
	if user_id != 0 {
		db := utils.GetDB()
		db.First(&myUserModel, user_id)
	}

	// 给上下文添加user信息字段
	c.Set("user_id", user_id)
	c.Set("my_user_model", myUserModel)
}

// 过滤不需要认证的路由
var no_auth_route_url = []string{
	"/api/login", "/api/register",
}

// Authorization Middleware:
// Extract auth_token in request cookie, and jwt authorization
func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// init context userid
		UpdateContextUserModel(c, 0)
		// access source list

		// access control
		for _, v := range no_auth_route_url {
			if c.Request.URL.Path == v {
				c.Next()
				return
			}
		}

		// token, err := request.ParseFromRequest(c.Request, MyAuth2Extractor, func(_ *jwt.Token) (interface{}, error) {
		// 	b := ([]byte(utils.SecretPassword))
		// 	return b, nil
		// })

		cookie, err := c.Cookie("auth_token")
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		token, err := jwt.Parse(cookie, func(_ *jwt.Token) (interface{}, error) {
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
			user_id := uint(claims["id"].(float64))
			UpdateContextUserModel(c, user_id)
		}
	}
}
