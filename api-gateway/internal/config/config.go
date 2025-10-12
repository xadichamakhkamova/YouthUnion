package config

import "github.com/spf13/viper"

type ServiceConfig struct {
	Host string
	Port int
}

type Config struct {
	ApiGateway  ServiceConfig
	UserService ServiceConfig
	TokenKey    string
	CertFile    string
	KeyFile     string
}

func Load(path string) (*Config, error) {

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		ApiGateway: ServiceConfig{
			Host: viper.GetString("api-gateway.host"),
			Port: viper.GetInt("api-gateway.port"),
		},
		UserService: ServiceConfig{
			Host: viper.GetString("services.user-service.host"),
			Port: viper.GetInt("services.user-service.port"),
		},

		TokenKey: viper.GetString("token.key"),

		CertFile: viper.GetString("file.cert"),
		KeyFile:  viper.GetString("file.key"),
	}
	return &cfg, nil
}
