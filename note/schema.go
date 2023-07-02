package note

type NoteInput struct {
	Content string `json:"content" binding:"required"`
}

type NoteQuery struct {
	Content string `form:"content"`
}
