package note

import (
	"fmt"

	"github.com/TheBigRoomXXL/gin-htmx-experiment/commons/db"
	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Markdown string
	Html     []byte
}

func Create(md string) (*Note, error) {
	maybeUnsafeHTML := markdown.ToHTML([]byte(md), nil, nil)
	html := bluemonday.UGCPolicy().SanitizeBytes(maybeUnsafeHTML)
	note := Note{Markdown: md, Html: html}
	err := db.Con.Create(&note).Error
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func Update(id int, md string) error {
	var note *Note
	err := db.Con.First(&note, id).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	maybeUnsafeHTML := markdown.ToHTML([]byte(md), nil, nil)
	note.Html = bluemonday.UGCPolicy().SanitizeBytes(maybeUnsafeHTML)
	note.Markdown = md
	db.Con.Updates(note)
	return nil
}

func Search(search string) (*[]Note, error) {
	var notes *[]Note
	query := db.Con.Select("*").
		Where("Markdown LIKE ?", "%"+search+"%").
		Limit(100).
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
