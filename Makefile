BASE_DOMAIN = "github.com/kaiobrito/repository-blogpost"
REPO = "memory"

run_todo:
	go run "${BASE_DOMAIN}/applications/todo/cmd"

run_todorepo:
	go run "${BASE_DOMAIN}/applications/todorepo/cmd" -repository $(REPO)

run_todogrpc:
	go run "${BASE_DOMAIN}/applications/todogrpc/cmd"

generate_proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		./applications/**/proto/*.proto