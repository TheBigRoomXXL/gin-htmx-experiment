package main

import (
	"github.com/TheBigRoomXXL/gin-htmx-experiment/commons/db"
	"github.com/TheBigRoomXXL/gin-htmx-experiment/note"
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
	router.Run("0.0.0.0:3001")
}
