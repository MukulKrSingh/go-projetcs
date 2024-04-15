package services

type NotesService struct {
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
	return data
}
