ps:
	@docker-compose ps

build:
	@go mod vendor
	@docker-compose build

start:
	@go mod vendor
	@docker-compose up --build -d

compose/up:
	@docker-compose up -d

exec:
	@docker exec -it go-readinglist-go-1 /bin/sh

test:
	@go test ./...

push:
	@go fmt ./...
	@cd frontend && yarn lint --fix
	@git add .
	@git commit -m ${MSG}
	@git push origin master

serve:
	@cd frontend && yarn lint --fix && yarn serve
