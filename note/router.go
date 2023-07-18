package note

import (
	"embed"
	"text/template"

	"github.com/gin-gonic/gin"
)

//go:embed templates
var templateFS embed.FS

func RegisterRoutes(router *gin.Engine) {
	router.Static("/assets", "./note/statics") //path relative to main.go
	router.GET("/", func(c *gin.Context) {
		t := template.Must(template.ParseFS(
			templateFS,
			"templates/index.html",
			"templates/header.html",
			"templates/search.html",
			"templates/create.html",
			"templates/notes.html",
		))

		notes, err := Search("")
		if err != nil {
			c.JSON(422, gin.H{"error": err})
			return
		}

		t.Execute(c.Writer, gin.H{"notes": notes})
	})

	router.GET("/notes", func(c *gin.Context) {
		t := template.Must(template.ParseFS(templateFS, "templates/notes.html"))

		var params NoteSearch
		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
			return
		}

		notes, err := Search(params.Search)
		if err != nil {
			c.JSON(422, gin.H{"error": err})
			return
		}

		t.Execute(c.Writer, gin.H{"notes": &notes})
	})

	router.POST("/notes", func(c *gin.Context) {
		t := template.Must(template.ParseFS(templateFS, "templates/notes.html"))

		var params NoteCreate
		if err := c.ShouldBind(&params); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
			return
		}

		note, err := Create(params.Content)
		if err != nil {
			c.JSON(422, gin.H{"error": err})
			return
		}

		t.Execute(c.Writer, gin.H{"notes": [1]*Note{note}})
	})
	router.PUT("/notes/:id", func(c *gin.Context) {
		var uri NoteUpdateUri
		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
			return
		}

		var body NoteUpdateBody
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
			return
		}

		err := Update(uri.Id, body.Content)
		if err != nil {
			c.JSON(422, gin.H{"error": err})
			return
		}
	})
}
