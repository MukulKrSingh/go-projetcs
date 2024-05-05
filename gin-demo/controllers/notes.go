package controllers

import (
	"gin-demo/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	notesService services.NotesService
}

func (n *NotesController) InitNotesController(notesService services.NotesService) *NotesController {
	n.notesService = notesService
	return n
}
func (n *NotesController) InitRoutes(router *gin.Engine) {
	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())
	notes.GET("/:id", n.GetNote())
	notes.POST("/", n.CreateNotes())
	notes.PUT("/", n.UpdateNote())
	notes.DELETE("/:id", n.DeleteNotes())
	// n.notesService = notesService
}
func (n *NotesController) GetNote() gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")
		noteId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		note, err := n.notesService.GetNoteService(noteId)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"note": note,
		})

	}
}

func (n *NotesController) GetNotes() gin.HandlerFunc {

	return func(c *gin.Context) {
		status := c.Query("status")
		var actualStatus bool
		var err error
		var all bool
		if status != "" {
			actualStatus, err = strconv.ParseBool(status)
			all = false
		} else {
			all = true
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		notes, err := n.notesService.GetNotesService(actualStatus, all)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"notes": notes,
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {

	type NoteBody struct {
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status"`
	}
	return func(c *gin.Context) {
		var noteBody NoteBody
		if err := c.BindJSON(&noteBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		note, err := n.notesService.CreateNotesService(
			noteBody.Title,
			noteBody.Status,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"notes": note,
		})
	}
}

func (n *NotesController) UpdateNote() gin.HandlerFunc {
	type NoteBody struct {
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status"`
		Id     int    `json:"id" binding:"required"`
	}
	return func(c *gin.Context) {
		var noteBody NoteBody
		if err := c.BindJSON(&noteBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		note, err := n.notesService.UpdateNotesService(
			noteBody.Id,
			noteBody.Title,
			noteBody.Status,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"notes": note,
		})
	}
}

func (n *NotesController) DeleteNotes() gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")
		noteId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = n.notesService.DeleteNotesService(
			noteId,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "note successfully deleted",
		})
	}
}
