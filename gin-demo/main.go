package main

import (
	"gin-demo/controllers"
	internal "gin-demo/internal/database"

	services "gin-demo/services"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	db := internal.InitDb()

	if db == nil {
		//Error while connecting to DB
		return
	}

	//Notes Services and Controllers
	notesService := &services.NotesService{}
	notesService.InitService(db)

	notesController := &controllers.NotesController{}
	notesController.InitNotesController(*notesService)
	notesController.InitRoutes(router)

	//Auth Services and controllers\
	authService := services.InitAuthService(db)

	authController := controllers.InitAuthController(authService)
	authController.InitRoutes(router)

	router.Run(":8080")
}
