package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type ServerConfiguration struct {
	Port                 string
	Secret               string
	LimitCountPerRequest int64
}

func ServerConfig() string {
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8080")

	bindAddress := fmt.Sprintf("%s:%s", viper.GetString("HOST"), viper.GetString("PORT"))
	log.Print("Server Running at ", bindAddress)

	return bindAddress
}