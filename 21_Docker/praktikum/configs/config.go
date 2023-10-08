package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	SERVER_PORT int
	DB_PORT     int
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func InitConfig() *Config {
	var conf = new(Config)
	conf = loadConfig()

	if conf == nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
	}

	return conf
}

func loadConfig() *Config {
	var conf = new(Config)

	err := godotenv.Load(".env")

	if err != nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}
	
	
	if value, found := os.LookupEnv("SERVER_PORT"); found {
		port, err := strconv.Atoi(value)
		if err != nil {
			logrus.Error("Config : invalid port value,", err.Error())
			return nil
		}
		conf.SERVER_PORT = port
	}

	if value, found := os.LookupEnv("DB_PORT"); found {
		port, err := strconv.Atoi(value)
		if err != nil {
			logrus.Error("Config : invalid port value,", err.Error())
			return nil
		}
		conf.DB_PORT = port
	}

	if value, found := os.LookupEnv("DB_HOST"); found {
		conf.DB_HOST = value
	}

	if value, found := os.LookupEnv("DB_USER"); found {
		conf.DB_USER = value
	}

	if value, found := os.LookupEnv("DB_PASSWORD"); found {
		conf.DB_PASSWORD = value
	}

	if value, found := os.LookupEnv("DB_NAME"); found {
		conf.DB_NAME = value
	}

	return conf


}