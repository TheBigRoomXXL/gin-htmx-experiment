package note

type NoteCreate struct {
	Content string `form:"content" binding:"required"`
}

type NoteSearch struct {
	Search string `form:"search"`
}
