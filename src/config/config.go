package config

import (
	"github.com/spf13/viper"
	"os"
	_const "social-work_notification_service/src/const"
)

type structure struct {
	App struct {
		Env         string `mapstructure:"env"`
		GrpcAddress string `mapstructure:"grpcAddress"`
	} `mapstructure:"app"`
}

type config struct {
	env         _const.Environment
	grpcAddress string
}

func GetConfig() (*config, error) {
	var result config

	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var structure structure
	err = viper.Unmarshal(&structure)
	if err != nil {
		return nil, err
	}

	result.env = _const.Environment(structure.App.Env)
	result.grpcAddress = structure.App.GrpcAddress

	return &result, nil

}
