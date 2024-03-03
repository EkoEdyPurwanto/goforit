# Create a set of database up/down migrations with a specific name.
migrate-create:
# -ext adalah file extension, artinya kita membuat file.sql || -dir adalah folder tempat simpan || usahakan jangan menggunakan spasi dalam penamaan file migration
	migrate create -ext sql -dir db/migrations ${APP_DB_NAME}

# Migrate database up
migrate-up:
	@echo "Migrating database up..."
	migrate -path db/migrations -database "postgresql://$(APP_DB_USER):$(APP_DB_PASSWORD)@$(APP_DB_HOST):$(APP_DB_PORT)/$(APP_DB_NAME)?sslmode=disable" -verbose up ${NN}
	@echo "Migration up completed successfully."
# Migrate database down
migrate-down:
	@echo "Migrating database down..."
	migrate -path db/migrations -database "postgresql://$(APP_DB_USER):$(APP_DB_PASSWORD)@$(APP_DB_HOST):$(APP_DB_PORT)/$(APP_DB_NAME)?sslmode=disable" -verbose down ${NN}
	@echo "Migration down completed successfully."

.PHONY: migrate-create migrate-up migrate-down