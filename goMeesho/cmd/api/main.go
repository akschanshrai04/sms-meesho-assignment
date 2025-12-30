package main

import (
	"goMeesho/config"
	"goMeesho/controllers"
	"goMeesho/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()

	r := gin.Default()
	r.Use(gin.Logger())

	sms := r.Group("/v0/user")
	sms.GET("/:phone/messages", controllers.GetSMSByPhone)

	r.Run(":8000")

}
