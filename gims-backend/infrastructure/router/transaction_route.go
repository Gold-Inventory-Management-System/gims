package router

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/controllers"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/repositories"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/usecases"
	"github.com/gin-gonic/gin"
)

func setUpTransactionRoute(r *gin.Engine) {
	goldRepository := repositories.NewGoldRepository(database.DB, database.GormDB)
	transactionRepository := repositories.NewTransactionRepository(database.DB, database.GormDB)
	transactionUsecase := usecases.NewTransactionUsecase(transactionRepository, goldRepository)
	transactionController := controllers.NewTransactionController(transactionUsecase)
	transaction := r.Group("/transaction")
	{
		transaction.POST("/new-buy-transaction", transactionController.NewTransactionBuy)
		transaction.POST("/new-sell-transaction", transactionController.NewTransactionSell)
		transaction.POST("/new-change-transaction", transactionController.NewTransactionChange)
		transaction.POST("/roll-back-transaction", transactionController.RollBackTransaction)
		transaction.GET("/get-all-transaction-join-gold", transactionController.GetAllTransactionJoinGold)
		transaction.GET("/get-transaction-by-transaction-type", transactionController.GetAllTransactionByTransactionType)
		transaction.GET("/get-transaction-by-time-interval/:range", transactionController.GetTransactionByTimeInterval)
		transaction.GET("/get-transaction-by-date/:date", transactionController.GetTransactionByDate)
		transaction.GET("/get-transaction-from-to", transactionController.GetTransactionFromTo)
		transaction.GET("get-daily-report", transactionController.GetDailyReport)
		transaction.GET("get-report-by-time-interval", transactionController.GetReportByTimeInterval)
		transaction.GET("get-dashboard", transactionController.GetDashboard)
	}
}