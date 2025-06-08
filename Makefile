APP_NAME=gurps-bchest-be
VERSION ?= 0.1.0
DOCKER_IMAGE=$(APP_NAME):$(VERSION)

default: build

-%:
	-@$(MAKE) $*

.PHONY: build
build:
	@echo "+ $@"
	docker build --target build -t $(DOCKER_IMAGE) .

.PHONY: test
test: build
	@echo "+ $@"
	go test ./internal/... -coverprofile=./cover.out -covermode=count
	go run github.com/vladopajic/go-test-coverage/v2 --config=./.testcoverage.yml

.PHONY: coverage
coverage: -test
	@echo "+ $@"
	go tool cover -html=cover.out -o=cover.html

.PHONY: package
package:
	@echo "+ $@"
	docker build --target bin -t $(DOCKER_IMAGE) .

.PHONY: run
run: package
	@echo "+ $@"
	docker compose up

.PHONY: clear-graphql
clear-graphql:
	@echo "+ $@"
	rm -f internal/graph/model/models_gen.go internal/graph/generated.go

.PHONY: graphql
graphql: clear-graphql
	@echo "+ $@"
	go run github.com/99designs/gqlgen generate

.PHONY: clear-mocks
clear-mocks:
	@echo "+ $@"
	@rm -rf internal/mocks/*

.PHONY: mocks
mocks: clear-mocks
	@echo "+ $@"
	docker run --rm -v .://src -w //src vektra/mockery:3
