package note

import (
	"net/http"

	"github.com/TheBigRoomXXL/note-api/db"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {
	// Validate input
	var input NoteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	note := NewNote(input)
	db.Con.Create(&note)
	c.JSON(http.StatusOK, note)
}

func QueryNote(c *gin.Context) { // Get model if exist
	query, err := prepareQuery(c)
	if err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	var notes []Note
	query.Find(&notes)

	c.JSON(http.StatusOK, notes)
}

func prepareQuery(c *gin.Context) (*gorm.DB, error) {
	// Prepare params
	var params NoteQuery
	if err := c.ShouldBindQuery(&params); err != nil {
		return nil, err
	}

	// Create select statement
	query := db.Con.Select("*")
	if params.Content != "" {
		query = query.Where("content LIKE ?", "%"+params.Content+"%")
	}

	// Add pagination
	paginate, err := db.GetPaginate(c)
	if err != nil {
		return nil, err
	}
	return query.Scopes(paginate), nil
}

func QueryNoteById(c *gin.Context) { // Get model if exist
	var note Note
	err := db.Con.First(&note, 100).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, note)
}
