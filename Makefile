.PHONY: all
all: compile

.PHONY: compile
compile:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o bin/canvas ./cmd/core/main.go

# .PHONY: docker-build
# docker-build:
# 	@docker build -t canvas .

.PHONY: docker-build
docker-build:
	@docker build -t canvas:local .

.PHONY: compose-start
compose-start:
	@docker-compose -f docker-compose.yaml up -d

.PHONY: compose-stop
compose-stop-local:
	@docker-compose -f docker-compose.yaml down

.PHONY: compose-remove
compose-remove-local:
	@docker-compose -f docker-compose.yaml rm -s -f

.PHONY: test
test:
	@go test -race -cover ./internal/core/...