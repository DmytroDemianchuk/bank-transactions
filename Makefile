run:
	docker-compose up -d --build

migrate-up:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up
migrate-down:
	migrate -path ./migrations -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable down

migrate-drop:
	migrate -path ./migrations -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable drop

swag:
	swag init -g cmd/main.go
