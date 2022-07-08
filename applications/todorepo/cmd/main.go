package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/kaiobrito/repository-blogpost/data"
	"github.com/kaiobrito/repository-blogpost/data/repository"
	"github.com/kaiobrito/repository-blogpost/external/todoapi"
)

type App struct {
	Repo repository.IRepository[data.Todo]
}

func main() {
	app := App{
		Repo: createTODOAPIRepository(),
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

func createTODOAPIRepository() repository.IRepository[data.Todo] {
	username := flag.String("username", "", "Username used to authenticate at Todo api")
	password := flag.String("password", "", "Password used to authenticate at Todo api")
	flag.Parse()

	return todoapi.CreateTODOAPIRepository(*username, *password)
}
