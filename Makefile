# Simple Makefile for a Go project
ifneq (,$(wildcard ./.env))
    include .env
    export
endif
# Build the application
args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

all: build

build:
	@echo "Building..."
	
	
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Create DB container
up:
	# @if docker compose up 2>/dev/null; then \
	# 	: ; \
	# else \
	# 	echo "Falling back to Docker Compose V1"; \
	# 	docker-compose up; \
	# fi
	@docker compose up

# Shutdown DB container
down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test $$(go list ./... | grep -v -e '/models' -e '/dtos' -e '/docs')

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/air-verse/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean

nah:
	@echo $(call args,defaultstring)
