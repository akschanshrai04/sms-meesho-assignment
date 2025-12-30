package main

import (
	"goMeesho/config"
	"goMeesho/database"
	"goMeesho/services"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()
	services.StartKafkaConsumer()
}
