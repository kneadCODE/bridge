COMPOSE_CMD := docker-compose -f docker-compose.yaml

test:
ifdef CI_ENV
	${COMPOSE_CMD} run api sh -c "go mod vendor"
endif
	${COMPOSE_CMD} run --rm api sh -c "go test -mod=vendor -coverprofile=c.out -failfast -timeout 5m ./..."

run-api:
	${COMPOSE_CMD} run --rm --service-ports api sh -c "go run -mod=vendor cmd/api/*.go"

teardown:
	${COMPOSE_CMD} down -v

vendor:
	${COMPOSE_CMD} run --rm api sh -c "go mod tidy && go mod vendor"
