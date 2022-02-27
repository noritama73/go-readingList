ps:
	@docker-compose ps

build:
	@go mod vendor
	@docker-compose build

start:
	@go mod vendor
	@docker-compose up --build -d

compose/up:
	@docker-compose up

exec:
	@docker exec -it go-readinglist-go-1 /bin/sh

test:
	@go test ./...

git:
	@go fmt ./...
	@cd frontend && yarn lint
	@git add .
	@git commit -m ${MSG}
	@git push origin master
