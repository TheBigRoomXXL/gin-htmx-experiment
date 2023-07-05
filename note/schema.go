package note

type NoteCreate struct {
	Content string `form:"content" binding:"required" validate:"nonzero"`
}

type NoteSearch struct {
	Search string `form:"search"`
}
