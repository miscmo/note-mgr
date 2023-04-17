package main

import (
	"github.com/gin-gonic/gin"
	"github.com/miscmo/note-mgr/api"
	"github.com/miscmo/note-mgr/repository"
	"github.com/miscmo/note-mgr/resource"
)

func main() {

	resource.InitDB()
	resource.InitMongo()
	repository.Init()

	r := gin.Default()
	//r.LoadHTMLGlob("templates/*.html")
	//r.Static("/static", "templates/static")

	v1 := r.Group("/v1")
	{
		// v1.GET("/", api.HomePage)
		// v1.GET("/health", api.Health)
		v1.POST("/save_note", api.SaveNote)
		v1.POST("/update_note", api.UpdateNote)
		v1.GET("/get_notes", api.GetNotes)
	}

	r.Run(":8090")
}
