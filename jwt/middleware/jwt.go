package middleware

import (
	"github.com/gin-gonic/gin"
	"godemos/jwt/e"
	"godemos/jwt/util"
	"net/http"
	"strings"
	"time"
)

var jwtSecret = ""

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := http.StatusOK
		code := 0

		// read token from query
		token := c.Query("token")
		if token == "" {
			// read token from header
			auth := c.GetHeader("Authorization")
			if len(auth) > 0 {
				token = strings.TrimPrefix(auth, "Bearer ")
			}
		}

		if token == "" {
			// read token from cookie
			cookie, err := c.Cookie("token")
			if err == nil && len(cookie) > 0 {
				token = cookie
			}
		}

		if token == "" {
			status = http.StatusBadRequest
			code = e.ErrorAuthTokenRequired
		} else {
			claims, err := util.ParseToken(token, jwtSecret)
			if err != nil {
				status = http.StatusUnauthorized
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				status = http.StatusUnauthorized
				code = e.ErrorAuthCheckTokenTimeout
			} else {
				c.Set("auth.id", claims.ID)
				c.Set("auth.username", claims.Username)
			}
		}

		if status != http.StatusOK {
			c.JSON(status, e.Error(code))

			c.Abort()
			return
		}

		c.Next()
	}
}
