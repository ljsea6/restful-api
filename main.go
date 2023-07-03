package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/ljsea6/restful-api/app/router"
	"github.com/ljsea6/restful-api/config"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	app.Run(":" + port)
}
