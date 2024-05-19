package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jwt/helper"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return
		}

		claim, err := helper.ValidateToken(clientToken)

		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claim.Email)
		c.Set("first_name", claim.FirstName)
		c.Set("last_name", claim.LastName)
		c.Set("uid", claim.Uid)
		c.Set("user_type", claim.User_type)

		c.Next()
	}

}
