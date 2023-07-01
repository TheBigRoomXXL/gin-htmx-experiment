package note

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Content string `json:"content"`
}

func NewNote(input NoteInput) *Note {
	note := new(Note)
	note.Content = input.Content
	return note
}
