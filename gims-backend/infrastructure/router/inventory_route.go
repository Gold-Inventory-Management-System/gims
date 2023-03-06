package router

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/controllers"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/middlewares"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/repositories"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/usecases"
	"github.com/gin-gonic/gin"
)

func setUpInventoryRoute(r *gin.Engine) {
	goldRepository := repositories.NewGoldRepository(database.DB, database.GormDB)
	goldUseCase := usecases.NewGoldUseCase(goldRepository)
	goldController := controllers.NewGoldController(goldUseCase)
	inventory := r.Group("/inventory")
	{
		inventory.GET("/find-gold-detail-by-code/:code", goldController.FindGoldDetailByCode)
		inventory.POST("/find-gold-detail-by-detail", goldController.FindGoldDetailByDetail)
		inventory.POST("/get-gold-detail-join-inventory-by-detail", goldController.GetGoldDetailJoinInventoryByDetail)
		inventory.GET("/get-front-gold", goldController.GetAllFrontGold)

		inventory.Use(middlewares.AuthorizeAdminOrOwner())
		inventory.POST("/new-gold", goldController.NewGold)
		inventory.POST("/add-gold", goldController.AddGold)
		inventory.GET("/get-all-gold-detail", goldController.GetAllGoldDetail)
		inventory.PUT("/edit-gold-detail", goldController.EditGoldDetail)
		inventory.GET("/get-all-detail-join-inventory", goldController.GetAllGoldDetailJoinInventory)
		inventory.PATCH("/delete-gold-detail", goldController.SetStatusGoldDetailToDelete)
		inventory.PATCH("/get-back-gold-detail/:id", goldController.SetStatusGoldDetailToNormal)
		inventory.PATCH("/set-gold-inventory-status", goldController.SetStatusGoldInventory)
		inventory.GET("/get-gold-detail-by-gold-detail-id/:id", goldController.GetGoldDetailByGoldDetailID)
		inventory.POST("/delete-gold-inventory-by-id", goldController.DeleteGoldInventoryByIDArray)
		inventory.PATCH("/set-tag-serial-number", goldController.SetTagSerialNumberGoldInventory)
		inventory.GET("/get-gold-by-tag-serial-number", goldController.GetGoldByTagSerailNumber)
		inventory.POST("/check-gold", goldController.CheckGold)
	}
}