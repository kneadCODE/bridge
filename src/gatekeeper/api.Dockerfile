FROM golang:1.19-alpine AS builder
WORKDIR /src/gatekeeper
COPY ../ /src
ENV CGO_ENABLED 0
ENV GOARCH amd64
ENV GOOS linux
RUN go mod vendor
RUN go clean -mod=vendor -i -cache ./... && go build -mod=vendor -o api cmd/api/*.go

FROM alpine:3
RUN apk update && \
   apk add ca-certificates && \
   update-ca-certificates && \
   rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=builder /src/gatekeeper/api /app/
ENTRYPOINT ./api
