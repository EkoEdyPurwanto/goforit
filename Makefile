# Create database migration
migrate-create:
# -ext adalah file extension, artinya kita membuat file.sql || -dir adalah folder tempat simpan || usahakan jangan menggunakan spasi dalam penamaan file migration
	migrate create -ext sql -dir db/migrations ${APP_DB_NAME}

# Migrate database up
migrate-up:
	migrate -path db/migrations -database "postgresql://$(APP_DB_USER):$(APP_DB_PASSWORD)@$(APP_DB_HOST):$(APP_DB_PORT)/$(APP_DB_NAME)?sslmode=disable" -verbose up

# Migrate database down
migrate-down:
	migrate -path db/migrations -database "postgresql://$(APP_DB_USER):$(APP_DB_PASSWORD)@$(APP_DB_HOST):$(APP_DB_PORT)/$(APP_DB_NAME)?sslmode=disable" -verbose down

.PHONY: migrate-create migrate-up migrate-down