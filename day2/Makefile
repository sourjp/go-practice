db-init:
	@docker run --name go-pg -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=todo -p 5432:5432 -d postgres:12.4
	@sleep 5
	@psql -h localhost -U root -W root -d todo -f todo.sql

db-clean:
	@docker kill go-pg; docker rm go-pg

db-create:
	@psql -h localhost -U root -W root -d todo -f todo.sql

db-access:
	@psql -h localhost -U root -W root -d todo

test:
	curl localhost:8080/api/v1/todos
	@echo 
	curl -X POST localhost:8080/api/v1/todos -d '{"title": "yesterday", "message": "cooked"}'
	@echo 
	curl -X POST localhost:8080/api/v1/todos -d '{"title": "today", "message": "cook"}'
	@echo 
	curl -X POST localhost:8080/api/v1/todos -d '{"title": "tomorrow", "message": "will cook"}'
	@echo 
	curl localhost:8080/api/v1/todos/3
	@echo 
	curl -X PUT localhost:8080/api/v1/todos -d '{"title": "tomorrow", "message": "will not cook"}'
	@echo 
	curl localhost:8080/api/v1/todos
	@echo 
	curl -X DELETE localhost:8080/api/v1/todos/1
	@echo 
	curl localhost:8080/api/v1/todos
	@echo 
	curl localhost:8080/api/v1/todos?limit=1
