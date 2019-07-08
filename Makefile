SCRIPT_DIR=./scripts

.PHONY: init
init:
	${SCRIPT_DIR}/init.sh

.PHONY: migrate
migrate:
	${SCRIPT_DIR}/migrate.sh