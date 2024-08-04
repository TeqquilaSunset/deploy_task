package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/imshawan/gin-backend-starter/services"
	"github.com/imshawan/gin-backend-starter/routers/middlewares"
	"github.com/imshawan/gin-backend-starter/models"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	router.POST("/sign-in", middlewares.ValidateRequestFields(&models.UserSigninRequest{}), services.SignIn)
}