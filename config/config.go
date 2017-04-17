package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Debug         bool
	ListeningPort int
	JWTSigningKey string
	DB            *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Username string
	Name     string
	Password string
}

func GetConfig() *Config {
	v := viper.New()

	v.AutomaticEnv()
	v.SetEnvPrefix("auth")

	loadDefaultSettings(v)

	return &Config{
		Debug:         v.GetBool("debug"),
		ListeningPort: v.GetInt("listening_port"),
		JWTSigningKey: v.GetString("jwt_signing_key"),
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     v.GetString("database_host"),
			Username: v.GetString("database_username"),
			Name:     v.GetString("database_name"),
			Password: v.GetString("database_password"),
		},
	}
}

func loadDefaultSettings(v *viper.Viper) {
	v.SetDefault("debug", false)
	v.SetDefault("listening_port", 3000)
	v.SetDefault("jwt_signing_key", "secret key to encrypt the token with")
	v.SetDefault("database_host", "localhost")
	v.SetDefault("database_username", "dbuser")
	v.SetDefault("database_name", "awesomeappname")
	v.SetDefault("database_password", "password")
}
