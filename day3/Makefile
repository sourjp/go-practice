db-init:
	@docker run --name go-pg -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=users -p 5432:5432 -d postgres:12.4
	@sleep 5
	@psql -h localhost -U root -W root -d users -f users.sql

db-clean:
	@docker kill go-pg; docker rm go-pg

db-create:
	@psql -h localhost -U root -W root -d users -f users.sql

db-access:
	@psql -h localhost -U root -W root -d users

test:
	curl -X POST localhost:8080/api/v1/users/ -d '{"name": "james", "passsword": "P@ssw0rd!"}'
	@echo 
	curl -X POST localhost:8080/api/v1/users/ -d '{"name": "john", "passsword": "P@ssw0rd!"}'
	@echo 
	curl -X POST localhost:8080/api/v1/users/ -d '{"name": "jane", "passsword": "P@ssw0rd!"}'
	@echo 
	curl localhost:8080/api/v1/users/3
	@echo 
	curl -X PUT localhost:8080/api/v1/users/3 -d '{"name": "Tom", "passsword": "P@ssw0rd!"}'
	@echo 
	curl localhost:8080/api/v1/users/3
	@echo 
	curl -X DELETE localhost:8080/api/v1/users/1
