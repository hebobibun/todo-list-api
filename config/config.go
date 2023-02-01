package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Mysql_Host     string
	Mysql_Port     int
	Mysql_User     string
	Mysql_Password string
	Mysql_DBName   string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}

	isRead := true

	if val, found := os.LookupEnv("MYSQL_HOST"); found {
		app.Mysql_Host = val
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_PORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.Mysql_Port = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_USER"); found {
		app.Mysql_User = val
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_PASSWORD"); found {
		app.Mysql_Password = val
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_DBNAME"); found {
		app.Mysql_DBName = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}
		err = viper.Unmarshal(&app)
		if err != nil {
			log.Println("error parse config : ", err.Error())
			return nil
		}

	}

	return &app
}
