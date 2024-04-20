package services

import (
	"fmt"
	internal "gin-demo/internal/models"

	"gorm.io/gorm"
)

type NotesService struct {
	db *gorm.DB
}

func (n *NotesService) InitService(databsase *gorm.DB) {
	n.db = databsase
	n.db.AutoMigrate(&internal.Notes{})
}

type Note struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (n *NotesService) GetNotesService() []Note {
	data := []Note{
		{
			Id:   1,
			Name: "Note 1",
		},
		{
			Id:   2,
			Name: "Note 2",
		},
	}

	return data

}

func (n *NotesService) CreateNotesService() Note {
	data := Note{
		Id:   3,
		Name: "Note 3",
	}

	err := n.db.Create(&internal.Notes{
		Id:     1,
		Title:  "Note-1",
		Status: true,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
	return data
}
