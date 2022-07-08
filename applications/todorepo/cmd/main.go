package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kaiobrito/repository-blogpost/data"
	"github.com/kaiobrito/repository-blogpost/data/repository"
)

type App struct {
	Repo repository.IRepository[data.Todo]
}

func main() {
	initialData := []data.Todo{}
	app := App{
		Repo: repository.CreateMemoryRepository(initialData),
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
