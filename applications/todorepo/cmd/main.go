package main

import (
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	gRepo "github.com/kaiobrito/repository-blogpost/applications/todogrpc/repository"
	"github.com/kaiobrito/repository-blogpost/data"
	"github.com/kaiobrito/repository-blogpost/data/repository"
	"github.com/kaiobrito/repository-blogpost/external/todoapi"
)

type App struct {
	Repo repository.IRepository[data.Todo]
}

func main() {
	app := App{
		Repo: createRepository(),
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
func createRepository() repository.IRepository[data.Todo] {
	repo := flag.String("repository", "", "Type of repository that will be used")
	flag.Parse()

	if *repo == "api" {
		log.Println("Using API Repository")
		return createTODOAPIRepositoryFromEnv()
	} else if *repo == "grpc" {
		log.Println("Using GRPC Repository")
		return gRepo.CreateTodoGRPCService("localhost:50051")
	}

	log.Println("Using memory Repository")
	initialData := []data.Todo{}
	return repository.CreateMemoryRepository(initialData)
}

func createTODOAPIRepositoryFromEnv() repository.IRepository[data.Todo] {
	token := os.Getenv("API_TOKEN")

	return todoapi.CreateTODOAPIRepositoryWithToken(token)
}
