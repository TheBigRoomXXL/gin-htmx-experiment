package note

type NoteInput struct {
	Content string `json:"content" binding:"required"`
}

type NoteQuery struct {
	ID      string `form:"id"`
	Content string `form:"content"`
}
