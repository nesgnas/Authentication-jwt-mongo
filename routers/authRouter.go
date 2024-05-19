package routers

import (
	"github.com/gin-gonic/gin"
	"jwt/controler"
)

func AuthRoutes(incomingRouter *gin.Engine) {
	incomingRouter.POST("user/signup", controler.SignUpUser())
	incomingRouter.POST("user/login", controler.LoginUser())

}
