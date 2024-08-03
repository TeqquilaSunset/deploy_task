package routers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imshawan/gin-backend-starter/helpers"
	"github.com/imshawan/gin-backend-starter/routers/middlewares"
	"github.com/spf13/viper"
)

func registerInternalRoutes(router *gin.Engine) {
	router.NoRoute(func(context *gin.Context) {
		helpers.FormatAPIResponse(context, http.StatusNotFound, errors.New("route not found"))
	})
	router.GET("/health", func(context *gin.Context) { helpers.FormatAPIResponse(context, http.StatusOK, gin.H{"health": "Ok"}) })
}

func SetupRouters() *gin.Engine {

	environment := viper.GetBool("DEBUG")
	if environment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	allowedHosts := viper.GetString("ALLOWED_HOSTS")
	router := gin.New()
	router.SetTrustedProxies([]string{allowedHosts})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())

	registerInternalRoutes(router) //Internal routes registeration

	return router
}
