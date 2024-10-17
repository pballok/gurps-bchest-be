APP_NAME=gurps-bchest-be
VERSION ?= 0.1.0
DOCKER_IMAGE=$(APP_NAME):$(VERSION)

default: build

.PHONY: build
build:
	@echo "+ $@"
	docker build --target build -t $(DOCKER_IMAGE) .

.PHONY: package
package: build
	@echo "+ $@"
	docker build --target bin -t $(DOCKER_IMAGE) .

.PHONY: run
run: package
	@echo "+ $@"
	docker run --rm -it --name $(APP_NAME) -p 8080:8080 $(DOCKER_IMAGE)
