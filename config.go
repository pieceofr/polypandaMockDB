package main

import (
	"log"

	"github.com/spf13/viper"
)

/*ServiceConfig Mock Service Config*/
type ServiceConfig struct {
	SQLEndpoint   string
	SQLUser       string
	SQLPwd        string
	SQLDB         string
	SQLPandaTable string
}

func initConfig() ServiceConfig {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Fatalln("Fatal error config file: %s \n", err)
	}
	config := ServiceConfig{SQLEndpoint: viper.GetString("database.sqlendpoint"),
		SQLUser:       viper.GetString("database.sqluser"),
		SQLPwd:        viper.GetString("database.sqlpwd"),
		SQLDB:         viper.GetString("database.polypandadb"),
		SQLPandaTable: viper.GetString("database.pandatable"),
	}
	return config
}
