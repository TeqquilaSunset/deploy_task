package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/imshawan/gin-backend-starter/models"
	"github.com/imshawan/gin-backend-starter/routers/middlewares"
	"github.com/imshawan/gin-backend-starter/services"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	router.GET("/", middlewares.IsAuthenticated(), services.UserProfile)
	router.POST("/register", middlewares.ValidateRequestFields(&models.UserRequest{}), services.RegisterUser)
}
