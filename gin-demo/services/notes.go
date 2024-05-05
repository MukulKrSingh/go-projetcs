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

func (n *NotesService) GetNotesService(status bool, all bool) ([]*internal.Notes, error) {

	var notes []*internal.Notes

	if !all {
		if err := n.db.Where("status = ?", status).Find(&notes).Error; err != nil {
			return nil, err
		}
	} else {
		if err := n.db.Find(&notes).Error; err != nil {
			return nil, err
		}
	}

	return notes, nil
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

func (n *NotesService) UpdateNotesService(id int, title string, status bool) (*internal.Notes, error) {

	var note *internal.Notes

	if err := n.db.Where("id = ? ", id).First(&note).Error; err != nil {
		return nil, err
	}

	note.Title = title
	note.Status = status

	if err := n.db.Save(&note).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	// note := &internal.Notes{
	// 	Id:     id,
	// 	Title:  title,
	// 	Status: status,
	// }
	// n.db.Update(fmt.Sprint(note.Id), note)

	return note, nil
}

func (n *NotesService) DeleteNotesService(id int64) error {

	var note *internal.Notes

	if err := n.db.Where("id = ? ", id).First(&note).Error; err != nil {
		return err
	}

	if err := n.db.Where("id = ?", id).Delete(&note).Error; err != nil {
		return err
	}

	// if err := n.db.Save(&note).Error; err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// note := &internal.Notes{
	// 	Id:     id,
	// 	Title:  title,
	// 	Status: status,
	// }
	// n.db.Update(fmt.Sprint(note.Id), note)

	return nil
}
