# Makefile is used to handle local development easily, without having to manage remembering multiple commands.
# Handles varios methods of building, and running with different profiles.
.PHONY: build run

all: build run
build:
	@go build -C src
run:
	@cd src && ./baby_names