package main

import (
	"github.com/briancolinger/go-crud/initializers"
	"github.com/briancolinger/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
