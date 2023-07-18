package note

type NoteCreate struct {
	Content string `form:"content" binding:"required" validate:"nonzero"`
}

type NoteUpdateUri struct {
	Id int `uri:"id" binding:"required"`
}

type NoteUpdateBody struct {
	Content string `form:"content" binding:"required" validate:"nonzero"`
}

type NoteSearch struct {
	Search string `form:"search"`
}
