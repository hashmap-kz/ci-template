FROM golang:1.24-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/root/go-build go mod download -x

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go test -v ./... && go build -ldflags="-s -w" -o ./ci-template

FROM alpine:latest AS build-release-stage

COPY --from=build-stage /app/ci-template /usr/local/bin/ci-template
RUN chmod +x /usr/local/bin/ci-template

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

ENTRYPOINT ["/usr/local/bin/ci-template"]
