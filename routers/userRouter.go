package routers

import (
	"github.com/gin-gonic/gin"
	"jwt/controler"
	"jwt/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("/user", controler.GetUsers())
	incomingRoutes.GET("/user/:user_id", controler.GetUser())

}
