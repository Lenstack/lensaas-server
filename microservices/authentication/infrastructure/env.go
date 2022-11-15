package infrastructure

import (
	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigFile("./microservices/authentication/.env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}
