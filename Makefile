COMPOSE_CMD := docker-compose -f build/docker-compose.yaml

teardown:
	${COMPOSE_CMD} down -v

# GoLib
golib-test:
ifdef CI_ENV
	${COMPOSE_CMD} run golib sh -c "go mod vendor"
endif
	${COMPOSE_CMD} run --rm golib sh -c "go test -mod=vendor -coverprofile=c.out -failfast -timeout 5m ./..."

golib-vendor:
	${COMPOSE_CMD} run --rm golib sh -c "go mod tidy && go mod vendor"

# Gatekeeper
gatekeeper-test:
ifdef CI_ENV
	${COMPOSE_CMD} run gatekeeper-api sh -c "go mod vendor"
endif
	${COMPOSE_CMD} run --rm gatekeeper-api sh -c "go test -mod=vendor -coverprofile=c.out -failfast -timeout 5m ./..."

gatekeeper-api:
	${COMPOSE_CMD} run --rm --service-ports gatekeeper-api sh -c "go run -mod=vendor cmd/api/*.go"

gatekeeper-vendor:
	${COMPOSE_CMD} run --rm gatekeeper-api sh -c "go mod tidy && go mod vendor"
