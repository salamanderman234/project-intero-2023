package config

import "github.com/spf13/viper"

func SetConfig(envUrl string) {
	viper.SetConfigFile(envUrl)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}