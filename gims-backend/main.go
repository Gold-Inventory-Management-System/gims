package main

import (
	"os"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.DB = database.NewDb()
	database.GormDB = database.NewGormDb()
	defer database.DB.Close()

	router.SetupRoutes(r)
	r.Run(os.Getenv("PORT"))
}