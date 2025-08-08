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

# Откат миграций с указанием количества шагов
migrate-down-steps:
	$(MIGRATE) down $(n)

# Принудительное выставление версии
migrate-force:
	$(MIGRATE) force $(v)


# Запуск приложения
run:
	go run cmd/main.go

# генерируем все
gen-all:
	make gen-tasks
	make gen-users

# генерируем таски
gen-tasks:
	oapi-codegen -config openapi/.openapi -include-tags Tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

# генерируем пользователей
gen-users:
	oapi-codegen \
	  -config openapi/.openapi \
	  -include-tags Users \
	  -package users \
	  openapi/openapi.yaml > internal/web/users/api.gen.go

lint:
	golangci-lint run --output.text.colors=true


.PHONY: run migrate migrate-down migrate-new