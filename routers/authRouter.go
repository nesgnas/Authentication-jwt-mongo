package routers

import (
	"github.com/gin-gonic/gin"
	"jwt/controler"
	"jwt/middleware"
)

func AuthRoutes(incomingRouter *gin.Engine) {
	incomingRouter.Use(middleware.CORSMiddleware())
	incomingRouter.POST("user/signup", controler.SignUpUser())
	incomingRouter.POST("user/login", controler.LoginUser())

}
