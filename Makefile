all: build

.PHONY: build
build:
	@docker build . --target bin
