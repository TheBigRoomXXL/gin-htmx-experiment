package note

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {

	notesRouter := router.Group("/notes")
	notesRouter.POST("", CreateNote)
	notesRouter.GET("", QueryNote)
	notesRouter.GET("/:id", QueryNoteById)
	return
}
