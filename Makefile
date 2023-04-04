# ----------------------------
# Env Variables
# ----------------------------
DOCKER_COMPOSE_FILE ?= docker-compose.yaml
DATABASE_CONTAINER ?= db
API_CONTAINER ?= api

# ----------------------------
# api Methods
# ----------------------------
api-setup: create-db db-migrate sqlboiler-psql
api-run:
	docker compose up -d ${API_CONTAINER}

# ----------------------------
# database Methods
# ----------------------------
## Start postgres database container only
create-db:
	docker compose up -d ${DATABASE_CONTAINER}

## Create a DB migration files e.g `make migrate-create name=test`
migrate-create:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate create -ext sql -dir /migrations $(name)

## Run migrations UP
db-migrate:
	docker compose --profile tools run --rm migrate up

## Rollback migrations against non test DB
db-redo:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate down

sqlboiler-psql:
	 cd api && sh -c "sqlboiler psql"
