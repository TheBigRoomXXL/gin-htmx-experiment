package note

import (
	"fmt"
	"net/http"

	"github.com/TheBigRoomXXL/note-api/db"

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
	db.DB.Create(&note)
	c.JSON(http.StatusOK, note)
}

func QueryNote(c *gin.Context) { // Get model if exist
	var params NoteQuery
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	var notes []Note
	query := db.DB.Select("*")
	if params.Content != "" {
		query = query.Where("content LIKE ?", "%"+params.Content+"%")
	}

	result := query.Find(&notes)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, notes)
}

func QueryNoteById(c *gin.Context) { // Get model if exist
	var note Note
	result := db.DB.First(&note, 10)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, note)
}
