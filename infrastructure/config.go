package infrastructure

import (
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to load up env configuration")
	}

}
