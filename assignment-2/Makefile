.PHONY: migration
migration:
	migrate create -ext sql -dir db/migrations/$(module) $(name)

.PHONY: migrate
migrate:
	migrate -path db/migrations/$(module) -database "postgres://postgresuser:postgrespassword@127.0.0.1:5432/postgres?sslmode=disable&search_path=public" -verbose up

.PHONY: rollback
rollback:
	migrate -path db/migrations/$(module) -database "postgres://postgresuser:postgrespassword@127.0.0.1:5432/postgres?sslmode=disable&search_path=public" -verbose down 1

.PHONY: rollback-all
rollback-all:
	migrate -path db/migrations/$(module) -database "postgres://postgresuser:postgrespassword@127.0.0.1:5432/postgres?sslmode=disable&search_path=public" -verbose down -all

.PHONY: force-migrate
force-migrate:
	migrate -path db/migrations/$(module) -database "postgres://postgresuser:postgrespassword@127.0.0.1:5432/postgres?sslmode=disable&search_path=public" -verbose force $(version)

.PHONY: compile-server
compile-server:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o deploy/hacktiv-assignment-2 main.go

.PHONY: docker-build-server
docker-build-server:
	docker build --no-cache -t hacktiv-assignment-2:latest -f Dockerfile .

.PHONY: run
run:
	docker-compose up

.PHONY: compile-build-run
compile-build-run: compile-server docker-build-server run