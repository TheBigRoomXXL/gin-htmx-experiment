package note

import (
	"html"

	"github.com/TheBigRoomXXL/note-api/db"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Content string
}

func Create(content string) (*Note, error) {
	escaped_content := html.EscapeString(content)
	note := Note{Content: escaped_content}
	err := db.Con.Create(&note).Error
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func Search(search string) (*[]Note, error) {
	var notes *[]Note
	query := db.Con.Select("*").
		Where("content LIKE ?", "%"+search+"%").
		Limit(10).
		Order("id DESC")

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
