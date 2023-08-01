package routes

import (
	controller "github.com/amit/go-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.SignUp())
	incomingRoutes.POST("users/login", controller.Login())
}
