package routes

import (
	controller "github.com/amit/go-jwt/controllers"
	"github.com/amit/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRouter(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/users", controller.GetUsers())
	incomingRoutes.POST("/users/:user_id", controller.GetUser())
}