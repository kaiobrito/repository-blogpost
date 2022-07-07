package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kaiobrito/repository-blogpost/data"
)

type App struct {
	Todos []data.Todo
}

func main() {
	app := App{
		Todos: []data.Todo{},
	}

	r := setupRouter(&app)
	r.Run()
}

func setupRouter(app *App) *gin.Engine {
	r := gin.Default()
	r.GET("todos/", app.GetTodos)
	r.GET("todos/:id", app.GetTodoById)
	r.POST("todos/:id", app.EditTodos)
	r.POST("todos/", app.CreateTodos)

	return r
}
