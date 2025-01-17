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

compose/down:
	@docker-compose down

exec:
	@docker exec -it gorl-backend /bin/sh

ci-test:
	@go mod vendor
	@PORT=8088 DRIVER=mysql DSN='root:root_test@tcp(127.0.0.1:3307)/gorl_db_test' MYSQL_ROOT_PASSWORD=root_test MYSQL_DATABASE=gorl_db_test TZ='Asia/Tokyo' go test ./... -coverpkg=./... $(shell go list ./... | grep -v 'vendor') -p 1

push:
	@go fmt ./...
	@cd frontend && yarn lint --fix
	@git add .
	@git commit -m ${MSG}
	@git push origin master

serve:
	@cd frontend && yarn lint --fix && yarn serve

cover:
	@go mod vendor
	@PORT=8088 DRIVER=mysql DSN='root:root_test@tcp(127.0.0.1:3307)/gorl_db_test' MYSQL_ROOT_PASSWORD=root_test MYSQL_DATABASE=gorl_db_test TZ='Asia/Tokyo' go test ./... -coverpkg=./... -coverprofile=cover.out.tmp $(shell go list ./... | grep -v 'vendor') -p 1
	@cat cover.out.tmp | grep -v "**_mock.go" | grep -v "wire_gen.go" > cover.out
	@rm cover.out.tmp
	@go tool cover -html=cover.out -o cover.html
	@explorer.exe cover.html

lint:
	@cd frontend && yarn lint --fix
