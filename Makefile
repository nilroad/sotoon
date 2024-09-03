COMPOSE_FILES=docker-compose.yml
COMPOSE_PROFILES=
COMPOSE_COMMAND=docker-compose

ifeq (, $(shell which $(COMPOSE_COMMAND)))
	COMPOSE_COMMAND=docker compose
	ifeq (, $(shell which $(COMPOSE_COMMAND)))
		$(error "No docker compose in path, consider installing docker on your machine.")
	endif
endif

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

ifeq ($(APP_ENV),develop)
	COMPOSE_FILES=docker-compose.yml -f docker-compose-dev.yml
endif

# If the first argument is "log"...
ifeq (log,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif


help:
	@echo "env"
	@echo "==> Create .env file"
	@echo ""
	@echo "init"
	@echo "==> init project"
	@echo "setup"
	@echo "==> setup project"
	@echo ""
	@echo "up"
	@echo "==> Create and start containers"
	@echo ""
	@echo "build-up"
	@echo "==> Create and build all containers"
	@echo ""
	@echo "status"
	@echo "==> Show currently running containers"
	@echo ""
	@echo "destroy"
	@echo "==> Down all the containers, keeping their data"
	@echo ""
	@echo "mysql-shell"
	@echo "==> Create an interactive shell for mysql"
	@echo "swagger-generate"
	@echo "==> Create swagger (OpenAPI) document"
	@echo ""
env:
	@[ -e ./.env ] || cp -v ./.env.example ./.env

init:
	@echo installing swag
	go install github.com/swaggo/swag/cmd/swag@v1.16.3
	@echo installing golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3

up:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) up -d

build-up:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) up --build -d

build-no-cache:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) build --no-cache

status:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) ps $(RUN_ARGS)

down:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) down --remove-orphans $(RUN_ARGS)

purge:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) down --remove-orphans --volumes $(RUN_ARGS)

mysql-shell:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) exec -u 0 mysql mysql -hmysql -u$(MYSQL_DATABASE_USERNAME) -D$(MYSQL_DATABASE_NAME) -p$(MYSQL_DATABASE_PASSWORD)

redis-shell:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) exec -u 0 redis redis-cli

.PHONY: log
log:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) logs -f $(RUN_ARGS)

generate-swagger:
	swag init -q --parseDependency --parseInternal -g router.go -d internal/api/rest

lint:
	golangci-lint run

build-clean:
	rm -rf ./build

build-arm: build-clean
	go generate ./...
	env GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -trimpath -v -a -ldflags "-w -s" \
            -o build/ ./cmd/...

build-linux: build-clean
	go generate ./...
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -v -a -ldflags "-w -s" \
			-o build/ ./cmd/...

.PHONY: log
shell:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) exec $(RUN_ARGS) bash

shell-as-root:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) exec -u 0 $(RUN_ARGS) bash

precommit-hook:
	@[ -f ./precommit-hook.sh ] && cp -v ./precommit-hook.sh ./.git/hooks/pre-commit && echo "precommit hook installed" || echo "error: could not find precommit-hook.sh"