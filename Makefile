BASE_DOMAIN = "github.com/kaiobrito/repository-blogpost"

run_todo:
	go run "${BASE_DOMAIN}/applications/todo/cmd"

run_todorepo:
	go run "${BASE_DOMAIN}/applications/todorepo/cmd"