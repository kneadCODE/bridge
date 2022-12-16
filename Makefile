COMPOSE_CMD := docker-compose -f build/docker-compose.yaml

teardown:
	${COMPOSE_CMD} down -v

# GoLib
ifdef CI_ENV
golib-test: golib-vendor
endif
golib-test:
	${COMPOSE_CMD} run --rm golib sh -c "go test -mod=vendor -coverprofile=c.out -failfast -timeout 5m ./..."

golib-vendor:
	${COMPOSE_CMD} run --rm golib sh -c "go mod vendor"

# Gatekeeper
ifdef CI_ENV
gatekeeper-test: gatekeeper-vendor
endif
gatekeeper-test:
	${COMPOSE_CMD} run --rm gatekeeper-api sh -c "go test -mod=vendor -coverprofile=c.out -failfast -timeout 5m ./..."

gatekeeper-api:
	${COMPOSE_CMD} run --rm --service-ports gatekeeper-api sh -c "go run -mod=vendor cmd/api/*.go"

gatekeeper-vendor:
	${COMPOSE_CMD} run --rm gatekeeper-api sh -c "go mod vendor"

# Catalog
ifdef CI_ENV
catalog-test: catalog-vendor catalog-go-generate
endif
catalog-test:
	${COMPOSE_CMD} run --rm catalog-api sh -c "go test -mod=vendor -coverprofile=c.out -failfast -timeout 5m ./..."

catalog-api:
	${COMPOSE_CMD} run --rm --service-ports catalog-api sh -c "go run -mod=vendor cmd/api/*.go"

catalog-vendor:
	${COMPOSE_CMD} run --rm catalog-api sh -c "go mod vendor"

catalog-go-generate:
	${COMPOSE_CMD} run --rm catalog-api sh -c "go generate ./..."
