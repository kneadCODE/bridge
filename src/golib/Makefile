COMPOSE_CMD := docker-compose -f docker-compose.yaml

test:
ifdef CI_ENV
	${COMPOSE_CMD} run api sh -c "go mod vendor"
endif
	${COMPOSE_CMD} run --rm api sh -c "go test -mod=vendor -coverprofile=c.out -failfast -timeout 5m ./..."

teardown:
	${COMPOSE_CMD} down -v

vendor:
	${COMPOSE_CMD} run --rm api sh -c "go mod tidy && go mod vendor"
