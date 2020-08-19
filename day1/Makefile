run-1:
	@go run main.go

db-init:
	@docker run --name go-pg -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=todo -p 5432:5432 -d postgres:12.4
	@sleep 5
	@psql -h localhost -U root -W root -d todo -f todo.sql

db-clean:
	@docker kill go-pg
	@docker rm go-pg

db-createtbl:
	@psql -h localhost -U root -W root -d todo -f todo.sql

db-access:
	@psql -h localhost -U root -W root -d todo

t-get:
	@curl localhost:8080/api/v1/getall

t-post:
	@curl -X POST localhost:8080/api/v1/create -d '{"title": "study2", "message": "study2"}'
