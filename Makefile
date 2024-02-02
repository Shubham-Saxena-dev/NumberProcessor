.DEFAULT_GOAL := help

# Make version
ifneq (,)
$(error This makefile requires GNU Make)
endif

ARGS = $(filter-out $@,$(MAKECMDGOALS))
MAKEFLAGS += --silent

.PHONY: up
up:
	docker-compose up -d --build --remove-orphans

.PHONY: start
start: ## Start the containers
	docker-compose start

.PHONY: stop
stop: ## Stop the containers
	docker-compose stop