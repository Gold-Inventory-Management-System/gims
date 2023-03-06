package router

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(middlewares.CORSMiddleware())
	setUpAuthRoute(r)

	r.Use(middlewares.AuthorizeJWT())
	setUpInventoryRoute(r)
	setUpTransactionRoute(r)
}