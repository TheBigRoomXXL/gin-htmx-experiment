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
	note.RegisterRoutes(router)

	// GO!
	router.Run("localhost:4000")
}
