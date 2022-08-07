package main

import (
	"github.com/gin-gonic/gin"
)

type Todo struct {
	Title       string
	Description string
}

var todos []Todo

func addTodo(title string, description string) []Todo {
	todo := Todo{}
	todo.Title = title
	todo.Description = description

	todos = append(todos, todo)

	return todos
}

func init() {
	todos = []Todo{
		{Title: "test1", Description: "test1"},
	}
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./index.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"todos": todos})
	})

	router.POST("/create", func(ctx *gin.Context) {
		title := ctx.PostForm("title")
		description := ctx.PostForm("description")
		addTodo(title, description)

		ctx.Redirect(301, "/")
	})

	router.Run()
}
