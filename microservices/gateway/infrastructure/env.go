package infrastructure

import (
	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}
