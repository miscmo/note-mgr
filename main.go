package main

import (
	"github.com/gin-gonic/gin"
	"github.com/miscmo/note-mgr/api"
	"github.com/miscmo/note-mgr/repository"
)

func main() {
	//dao.InitDB()

	//node := dao.Node{
	//	Id: 123,
	//	Name: "test1",
	//	Type: 1,
	//	//Content: []dao.Note{
	//	//	dao.Note{
	//	//		Id: 2,
	//	//		Name: "test_note1",
	//	//		Content: "test",
	//	//	},
	//	//	dao.Note{
	//	//		Id: 3,
	//	//		Name: "test_note2",
	//	//		Content: "test",
	//	//	},
	//	//},
	//}
	//
	//dao.InsertNode(&node)

	repository.InitDB()

	r := gin.Default()
	//r.LoadHTMLGlob("templates/*.html")
	//r.Static("/static", "templates/static")

	v1 := r.Group("/v1")
	{
		// v1.GET("/", api.HomePage)
		// v1.GET("/health", api.Health)
		v1.POST("/save_note", api.SaveNote)
		// v1.POST("/addNodes", api.AddNodes)
		// v1.POST("/calc", api.Calc)
		// v1.POST("/updateNode", api.UpdateNode)
		// v1.POST("/addSubNode", api.AddSubNode)
	}

	r.Run(":8090")
}
