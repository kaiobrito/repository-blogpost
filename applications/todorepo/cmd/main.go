package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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
	token := flag.String("token", "", "Token used to authenticate at Todo api")
	flag.Parse()

	fmt.Println(token, username, password)

	if *token != "" {
		log.Println("Creating Repository with token")
		return todoapi.CreateTODOAPIRepositoryWithToken(*token)
	}

	if *username != "" && *password != "" {
		log.Println("Creating Repository with username and password")
		return todoapi.CreateTODOAPIRepository(*username, *password)
	}

	log.Println("Creating Repository from env variable")
	return createTODOAPIRepositoryFromEnv()
}

func createTODOAPIRepositoryFromEnv() repository.IRepository[data.Todo] {
	token := os.Getenv("API_TOKEN")

	return todoapi.CreateTODOAPIRepositoryWithToken(token)
}
