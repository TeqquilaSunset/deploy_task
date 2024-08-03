package main

import (
	"log"
	"time"
	"os"

	"github.com/imshawan/gin-backend-starter/configs"
	"github.com/imshawan/gin-backend-starter/routers"
	"github.com/spf13/viper"
)

func main() {
	// Pre-start, setup intial configurations before the server starts up
	if err := configs.SetupConfig(); err != nil {
		log.Printf("Error while setting up server config, %s", err)
	}

	viper.SetDefault("SERVER_TIMEZONE", "Asia/Kolkata")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	router := routers.SetupRouters()

	// Run the server with the configurations
	err := router.Run(configs.ServerConfig())

	if err != nil {
		log.Printf("Error while starting up the server, %s", err)
		os.Exit(1)
	}
}