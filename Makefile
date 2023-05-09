PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin
$(shell [ -f bin ] || mkdir -p $(PROJECT_BIN))
PATH := $(PROJECT_BIN):$(PATH)
GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint

.PHONY: run
run:
	docker-compose build --no-cache
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose stop


.PHONY: kill
kill:
	docker-compose rm -f
	docker-compose down --volumes && docker network prune --force

.PHONY: log
log:
	docker-compose logs -f app

.PHONY: install-linter
install-linter:
	### INSTALL GOLANGCI-LINT ###
	[ -f $(PROJECT_BIN)/golangci-lint ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) v1.46.2

.PHONY: lint
lint:
	### RUN GOLANGCI-LINT ###
	$(GOLANGCI_LINT) run ./... --config=./.golangci.yml

.PHONY: lint-fast
lint-fast:
	$(GOLANGCI_LINT) run ./... --fast --config=./.golangci.yml
