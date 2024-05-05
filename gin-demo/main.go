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

	notesService := &services.NotesService{}
	notesService.InitService(db)

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message":  "Hello World",
	// 		"test_var": "Welcome",
	// 	})
	// })

	// router.GET("/user/:name", func(ctx *gin.Context) {
	// 	name := ctx.Params.ByName("name")

	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"username": name,
	// 		"messgae":  "welcome",
	// 	})

	// })

	// router.POST("/update-user", func(ctx *gin.Context) {
	// 	type UserCred struct {
	// 		Email    string `json:"email" binding:"required"`
	// 		Password string `json:"password"`
	// 	}

	// 	user := UserCred{}
	// 	err := ctx.BindJSON(&user)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{
	// 			"status":  "error",
	// 			"message": err.Error(),
	// 		})
	// 	} else {

	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"Email":    user.Email,
	// 			"Message":  user.Password,
	// 			"messsage": "Welcome!!",
	// 		})
	// 	}
	// })

	// router.PUT("/update-user", func(ctx *gin.Context) {
	// 	type UserCred struct {
	// 		Email    string `json:"email" binding:"required"`
	// 		Password string `json:"password"`
	// 	}

	// 	user := UserCred{}
	// 	err := ctx.BindJSON(&user)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{
	// 			"status":  "error",
	// 			"message": err.Error(),
	// 		})
	// 	} else {

	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"Email":    user.Email,
	// 			"Message":  user.Password,
	// 			"messsage": "Welcome!! using PUT",
	// 		})
	// 	}
	// })
	// router.PATCH("/update-user", func(ctx *gin.Context) {
	// 	type UserCred struct {
	// 		Email    string `json:"email" binding:"required"`
	// 		Password string `json:"password"`
	// 	}

	// 	user := UserCred{}
	// 	err := ctx.BindJSON(&user)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{
	// 			"status":  "error",
	// 			"message": err.Error(),
	// 		})
	// 	} else {

	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"Email":    user.Email,
	// 			"Message":  user.Password,
	// 			"messsage": "Welcome!! using PATCH",
	// 		})
	// 	}
	// })
	// router.DELETE("/update-user", func(ctx *gin.Context) {
	// 	type UserCred struct {
	// 		Email    string `json:"email" binding:"required"`
	// 		Password string `json:"password"`
	// 	}

	// 	user := UserCred{}
	// 	err := ctx.BindJSON(&user)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{
	// 			"status":  "error",
	// 			"message": err.Error(),
	// 		})
	// 	} else {

	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"Email":    user.Email,
	// 			"Message":  user.Password,
	// 			"messsage": "Welcome!! using PUT",
	// 		})
	// 	}
	// })

	notesController := &controllers.NotesController{}

	notesController.InitNotesController(*notesService)
	notesController.InitRoutes(router)

	router.Run(":8080")
}
