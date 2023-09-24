default: help

.PHONY: help
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

.PHONY: build
build: # dockerized compile for the project
	docker-compose -f .docker/compose.yml build --no-cache

.PHONY: up
up: build # dockerized execution of the file
	docker-compose -f .docker/compose.yml up --remove-orphans
