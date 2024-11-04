APP_NAME=gurps-bchest-be
VERSION ?= 0.1.0
DOCKER_IMAGE=$(APP_NAME):$(VERSION)

default: build

.PHONY: build
build:
	@echo "+ $@"
	docker build --target build -t $(DOCKER_IMAGE) .

.PHONY: test
test:
	@echo "+ $@"
	go test -coverprofile=coverage.out ./internal/...

.PHONY: package
package: build
	@echo "+ $@"
	docker build --target bin -t $(DOCKER_IMAGE) .

.PHONY: run
run: package
	@echo "+ $@"
	docker run --rm --name $(APP_NAME) -p 8080:8080 $(DOCKER_IMAGE)

.PHONY: clear-graphql
clear-graphql:
	@echo "+ $@"
	rm -f internal/graph/model/models_gen.go internal/graph/generated.go

.PHONY: graphql
graphql: clear-graphql
	@echo "+ $@"
	go run github.com/99designs/gqlgen generate
