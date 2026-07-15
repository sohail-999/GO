package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	//AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	//Println(viper.GetString("database.name"))
	fmt.Println(viper.GetString("database.host"))
	fmt.Println(viper.GetInt("database.port"))

	fmt.Println(viper.GetInt("server.port"))
	//.Println(viper.GetString("database.user"))

}
