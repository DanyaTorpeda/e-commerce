package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Server `yaml:"server"`
	DB     `yaml:"db"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DB struct {
	User     string `yaml:"user"`
	Password string
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

func LoadConfig(logger *logrus.Logger) (*Config, error) {
	if err := godotenv.Load(); err != nil {
		logger.Debug("error occured loading .env file")
		return nil, err
	}

	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.BindEnv("DB_PASSWORD")

	if err := viper.ReadInConfig(); err != nil {
		logger.Debug("error occured reading configuration file")
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		logger.Debug("error occured unmarshaling configuration file")
		return nil, err
	}

	config.DB.Password = viper.GetString("DB_PASSWORD")

	return &config, nil
}
