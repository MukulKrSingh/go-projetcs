package controllers

import (
	"gin-demo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	notesService services.NotesService
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine, notesService services.NotesService) {
	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNotes())
	n.notesService = notesService
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"notes": n.notesService.GetNotesService(),
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"notes": n.notesService.CreateNotesService(),
		})
	}
}
