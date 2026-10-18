.PHONY: install-libpq employees-db run-k6 run-reader mockery-gen

-include .env
export DATABASE_URL

run-reader:
	go run ./cmd/reader
	
run-tests:
	go test -v ./internal/...

employees-db:
	brew install libpq wget
	wget https://raw.githubusercontent.com/neondatabase/postgres-sample-dbs/main/employees.sql.gz
	pg_restore -d "$$DATABASE_URL" -Fc employees.sql.gz --no-owner --no-privileges
	rm employees.sql.gz

employees-schema:
	brew install libpq
	pg_dump "$$DATABASE_URL" --schema-only --schema=employees --no-owner --no-privileges --file=data/schema/employees.sql
	
lego-db:
	brew install libpq wget
	wget https://raw.githubusercontent.com/neondatabase/postgres-sample-dbs/main/lego.sql
	psql -d "$$DATABASE_URL" -f lego.sql
	rm lego.sql

lego-schema:
	brew install libpq
	pg_dump "$$DATABASE_URL" --schema-only --schema=public --no-owner --no-privileges --file=data/schema/lego.sql

DOCKER_MIGRATE = docker run --rm --network host -v $(CURDIR)/data/migrations:/migrations migrate/migrate:v4.18.1

migrate-up:
	$(DOCKER_MIGRATE) -path=/migrations -database '$(DATABASE_URL)' up

migrate-down:
	$(DOCKER_MIGRATE) -path=/migrations -database '$(DATABASE_URL)' down 1

sqlc-gen:
	docker run --rm -v $(PWD):/src -w /src sqlc/sqlc:1.31.1 generate

mockery-gen:
	docker run --rm -e GOTOOLCHAIN=auto -v "$(PWD):/src" -w /src vektra/mockery:v3.8.0 --config .mockery.yml

K6_SCENARIO = tests/load/$(SCENARIO).test.js

run-k6:
	@test -n "$(SCENARIO)" || { echo "Usage: make run-k6 SCENARIO=<scenario> (e.g. SCENARIO=simple-index)" >&2; exit 1; }
	@test -f "$(K6_SCENARIO)" || { echo "Scenario not found: $(K6_SCENARIO)" >&2; exit 1; }
	@command -v k6 >/dev/null 2>&1 || { echo "k6 is required: https://grafana.com/docs/learning-paths/run-first-k6-test/install-k6" >&2; exit 1; }
	k6 run "$(K6_SCENARIO)"
