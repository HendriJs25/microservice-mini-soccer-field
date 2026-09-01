SERVICE ?=
VERSION ?=

check-version:
	@test -n "$(VERSION)" || (echo "VERSION is required. Example: make migrate-force VERSION=1" && exit 1)

check-service:
	@test -n "$(SERVICE)" || (echo "SERVICE is required. Example: make deploy SERVICE=user-service" && exit 1)

setup: check-service
	docker compose build $(SERVICE)
	docker compose up -d postgres
	docker compose run --rm $(SERVICE) migrate up
	docker compose run --rm $(SERVICE) seed
	docker compose up -d $(SERVICE)

deploy: check-service
	docker compose build $(SERVICE)
	docker compose up -d postgres
	docker compose run --rm $(SERVICE) migrate up
	docker compose up -d $(SERVICE)

migrate-up: check-service
	docker compose run --rm $(SERVICE) migrate up

migrate-down: check-service
	docker compose run --rm $(SERVICE) migrate down

migrate-force: check-version check-service
	docker compose run --rm $(SERVICE) migrate force -- $(VERSION)

migrate-version: check-service
	docker compose run --rm $(SERVICE) migrate version

seed: check-service
	docker compose run --rm $(SERVICE) seed

up: check-service
	docker compose up -d $(SERVICE)

build: check-service
	docker compose build $(SERVICE)

down:
	docker compose down