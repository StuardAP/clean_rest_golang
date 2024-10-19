# #!make
include .env


#!COLORS
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[0;33m
RESET=\033[0m

#! ENVIRONMENT VARIABLES
APP_NAME = go-api
BINARY_PATH = bin/$(APP_NAME)
MAIN_FILE = cmd/api/main.go

MIGRATION_DIR = migrations/
DATABASE_URL = "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?multiStatements=true"

#! DEFAULT COMMANDS

.PHONY: help
help:
	@printf "$(YELLOW) Develoment commands :$(RESET) \n"
	@echo "  build        Build the Go binary"
	@echo "  test         Run tests"
	@echo "  run          Run the application"
	@echo "  air          Run with Air (live reload)"
	@echo "  clean        Remove binary files"
	@printf "$(YELLOW) Migration commands :$(RESET) \n"
	@echo "  migrate-create NAME=your_migration_name"
	@echo "  migrate-up"
	@echo "  migrate-down"
	@echo "  migrate-goto VERSION=version_number"
	@echo "  migrate-force VERSION=version_number"
	@echo "  migrate-drop"
	@echo "  migrate-version"

#! BUILD THE GO BINARY
.PHONY: build
build:
	@go build -o $(BINARY_PATH) $(MAIN_FILE)
	@printf "$(GREEN) Build Complete: $(BINARY_PATH) $(RESET)\n"

#! RUN ALL TESTS
.PHONY: test
test:
	@go test -v ./...

#! RUN THE APPLICATION WITHOUT BUILDING THE BINARY
.PHONY: run
run:
	@go run $(MAIN_FILE)

#! RUN THE APPLICATION WITH AIR FOR LIVE RELOAD
.PHONY: air
air:
	@air

#! CLEAN THE BINARY AND TEMP FILES
.PHONY: clean
clean:
	@rm -rf $(BINARY_PATH)
	@printf "$(GREEN) Cleaned Up Binary Files. $(RESET)\n"


#################! MIGRATION COMMANDS #################

#! COMMAND TO CREATE A NEW MIGRATION FILE
.PHONY: migrate-create
migrate-create:
	@if [ -z "$(NAME)" ]; then \
		printf "$(RED) ERROR: Please Specify a NAME for the Migration.$(RESET) \n"; \
		exit 1; \
	fi
	@migrate create -ext sql -dir $(MIGRATION_DIR) $(NAME)
	@printf "$(GREEN) Migration Created: $(NAME) $(RESET)\n"

#! COMMAND TO APPLY ALL PENDING MIGRATIONS
.PHONY: migrate-up
migrate-up:
	@migrate -path $(MIGRATION_DIR) -database $(DATABASE_URL) up

#! COMMAND TO REVERT ALL MIGRATIONS
.PHONY: migrate-down
migrate-down:
	@migrate -path $(MIGRATION_DIR) -database $(DATABASE_URL) down

#! COMMAND TO MIGRATE TO A SPECIFIC VERSION
.PHONY: migrate-goto
migrate-goto:
	@if [ -z "$(VERSION)" ]; then \
		printf "$(RED)ERROR: Please specify a VERSION to migrate to.$(RESET)"; \
		exit 1; \
	fi
	@migrate -path $(MIGRATION_DIR) -database $(DATABASE_URL) goto $(VERSION)

#! COMMAND TO FORCE THE MIGRATION VERSION (IGNORING DIRTY STATE)
.PHONY: migrate-force
migrate-force:
	@if [ -z "$(VERSION)" ]; then \
		 printf "$(RED)ERROR: Please specify a VERSION to force.$(RESET)"; \
		exit 1; \
	fi
	@migrate -path $(MIGRATION_DIR) -database $(DATABASE_URL) force $(VERSION)

#! COMMAND TO DROP ALL MIGRATIONS
.PHONY: migrate-drop
migrate-drop:
	@migrate -path $(MIGRATION_DIR) -database $(DATABASE_URL) drop

#! COMMAND TO SHOW CURRENT MIGRATION VERSION
.PHONY: migrate-version
migrate-version:
	@migrate -path $(MIGRATION_DIR) -database $(DATABASE_URL) version

