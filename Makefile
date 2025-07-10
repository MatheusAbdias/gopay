-include .env

MIGRATE_BIN := migrate
MIGRATE_DIR := internal/database/migrations
DB_URL := postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable

.PHONY: create-migration migrate-up migrate-down

create-migration:
	@if [ -z "$(name)" ]; then \
		echo "Error: 'name' not set. Usage: make create-migration name=<migration_name>"; \
		exit 1; \
	fi
	$(MIGRATE_BIN) create -ext sql -dir $(MIGRATE_DIR) -seq $(name)

migrate-up:
	$(MIGRATE_BIN) -path $(MIGRATE_DIR) -database "$(DB_URL)" -verbose up


migrate-down:
	$(MIGRATE_BIN) -path $(MIGRATE_DIR) -database "$(DB_URL)" -verbose down
