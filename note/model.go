package note

import (
	"github.com/TheBigRoomXXL/note-api/db"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Content string `json:"content"`
}

func Create(content string) *Note {
	note := new(Note)
	note.Content = content
	db.Con.Create(&note)
	return note
}

func Search(search string) ([]*Note, error) {
	var notes []*Note
	query := db.Con.Select("*").Where("content LIKE ?", "%"+search+"%")
	err := query.Find(&notes).Error
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func Get(id uint) (*Note, error) {
	var note *Note
	err := db.Con.First(&note, id).Error
	if err != nil {
		return nil, err
	}
	return note, nil
}
