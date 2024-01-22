# Makefile

.PHONY: build-images
build-images:
	docker-compose build

.PHONY: all-up
all-up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: restart
restart-containers: down all-up

.PHONY: clean
clean:
	docker-compose down -v --remove-orphans
