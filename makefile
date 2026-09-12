.PHONY: dev infra up down db sqlc test lint

infra:
	docker compose up -d

down:
	docker compose down

db:
	docker compose up -d postgres

sqlc:
	sqlc generate

test:
	go test ./...

test-race:
	go test -race ./...

api:
	go run ./cmd/api

web:
	cd apps/web && pnpm dev
