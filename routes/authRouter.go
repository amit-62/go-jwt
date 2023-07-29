package routes

import (
	controller "go-jwt/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(incomingRoutes *gin.Engine){
	incomingRoutes.POST("users/signup", controller.SignUP())
	incomingRoutes.POST("users/login", controller.Login())
}