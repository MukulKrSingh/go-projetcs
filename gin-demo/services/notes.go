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

func (n *NotesService) CreateNotesService(title string, status bool) (*internal.Notes, error) {

	note := &internal.Notes{

		Title:  title,
		Status: status,
	}
	if err := n.db.Create(note).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return note, nil
}
