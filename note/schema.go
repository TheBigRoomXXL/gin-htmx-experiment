package note

type NoteCreate struct {
	Content string `json:"content" binding:"required"`
}

type NoteSearch struct {
	Search string `form:"search"`
}
