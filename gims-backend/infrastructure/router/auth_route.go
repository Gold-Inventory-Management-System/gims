package router

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/controllers"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/middlewares"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/repositories"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/usecases"
	"github.com/gin-gonic/gin"
)

func setUpAuthRoute(r *gin.Engine) {
	userRepository := repositories.NewUserRepository(database.DB, database.GormDB)
	userUseCase := usecases.NewUserUseCase(userRepository)
	userController := controllers.NewUserController(userUseCase)
	auth := r.Group("/auth")
	{
		auth.POST("/register-admin", userController.RegisterAdmin)
		auth.POST("/signin", userController.SignIn)

		auth.Use(middlewares.AuthorizeJWT())
		auth.POST("/register", userController.Register)
		auth.GET("/profile", userController.TestJWT)
		auth.PUT("/rename-username", userController.RenameUsername)
		auth.PUT("/update-password", userController.UpdatePassword)
		auth.GET("/query-all-user", userController.QueryAllUser)
		auth.DELETE("/delete-user/:username", userController.DeleteUser)
	}
}