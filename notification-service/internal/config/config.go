package config

import "github.com/spf13/viper"

type ServiceConfig struct {
	Host string
	Port int
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Config struct {
	Service   ServiceConfig
	WebSocket ServiceConfig
	Postgres  PostgresConfig
	CertFile  string
	KeyFile   string
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		Service: ServiceConfig{
			Host: viper.GetString("service.grpc.host"),
			Port: viper.GetInt("service.grpc.port"),
		},
		WebSocket:  ServiceConfig{
			Host: viper.GetString("service.websocket.host"),
			Port: viper.GetInt("service.websocket.port"),
		},
		Postgres: PostgresConfig{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetInt("postgres.port"),
			Database: viper.GetString("postgres.dbname"),
			User:     viper.GetString("postgres.user"),
			Password: viper.GetString("postgres.password"),
		},

		CertFile: viper.GetString("file.cert"),
		KeyFile:  viper.GetString("file.key"),
	}

	return &cfg, nil
}
