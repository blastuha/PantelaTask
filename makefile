# переменные
DB_DSN := "postgres://postgres:aZAz1998@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# Запуск приложения
run:
	go run cmd/main.go

gen:
	oapi-codegen -config openapi/.openapi -include-tags Tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go


.PHONY: run migrate migrate-down migrate-new