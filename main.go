package main

import (
	"github.com/TheBigRoomXXL/note-api/db"
	"github.com/TheBigRoomXXL/note-api/note"
	"github.com/gin-gonic/gin"
)

func main() {
	//Init DB
	db.ConnectDatabase()
	err := db.Con.AutoMigrate(&note.Note{})
	if err != nil {
		return
	}

	//Init Router
	router := gin.Default()
	router.GET("/notes", note.QueryNote)
	router.GET("/notes/:id", note.QueryNoteById)
	router.POST("/notes", note.CreateNote)

	// GO!
	router.Run("localhost:8080")
}
