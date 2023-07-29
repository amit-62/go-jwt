package routes

import (
	controller "go-jwt/controllers"
	"go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRouter(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/users", controller.GetUsers())
	incomingRoutes.POST("/users/:user_id", controller.GetUser())
}