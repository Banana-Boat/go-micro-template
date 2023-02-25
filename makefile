MIGRATE_DB_URL=mysql://root:12345@tcp(localhost:3306)/test

mysql:
	docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=12345 -e MYSQL_DATABASE=test -d mysql:8.0
	
migrate_init:
	rm -rf internal/db/migration/** && migrate create -ext sql -dir internal/db/migration -seq init_schema

migrate_up:
	migrate -path internal/db/migration -database "${MIGRATE_DB_URL}" -verbose up

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

docker_build:
	docker build -t test-server:latest .

.PHONY: mysql migrate_init migrate_up sqlc test server docker_build
