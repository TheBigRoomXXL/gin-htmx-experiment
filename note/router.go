package note

import (
	"embed"
	"text/template"

	"github.com/gin-gonic/gin"
)

//go:embed templates
var templateFS embed.FS

func RegisterRoutes(router *gin.Engine) {

	router.StaticFile("/favicon.ico", "note/statics/favicon.ico")
	router.StaticFile("/style.css", "note/statics/style.css")
	router.GET("/", func(c *gin.Context) {
		t := template.Must(template.ParseFS(
			templateFS,
			"templates/index.html",
			"templates/search.html",
			"templates/notes.html",
		))

		var notes []*Note
		var err error
		if notes, err = Search(""); err != nil {
			c.JSON(422, gin.H{"error": err})
			return
		}

		t.Execute(c.Writer, gin.H{"notes": &notes})
	})

	router.GET("/notes", func(c *gin.Context) {
		t := template.Must(template.ParseFS(templateFS, "templates/notes.html"))

		// Parse input
		var params NoteSearch
		c.BindQuery(&params)

		var notes []*Note
		var err error
		if notes, err = Search(params.Search); err != nil {
			c.JSON(422, gin.H{"error": err})
			return
		}

		t.Execute(c.Writer, gin.H{"notes": notes})
	})

	router.POST("/notes", func(c *gin.Context) {
		t := template.Must(template.ParseFS(templateFS, "templates/notes.html"))

		// Parse input
		var params NoteCreate
		c.BindQuery(&params)

		note := Create(params.Content)

		t.Execute(c.Writer, gin.H{"notes": [1]*Note{note}})
	})
}
