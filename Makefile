GO ?= go
DOCKER ?= docker

.PHONY: test
test: generate-mocks
	$(GO) test --race ./...
.PHONY: generate-mocks
generate-mocks:
	$(DOCKER) run -v "$(PWD)":/app -w /app vektra/mockery
.PHONY: hash-migrations
hash-migrations:
	$(DOCKER) run --rm --net=host \
	-v $(PWD)/internal/migrations:/migrations \
	arigaio/atlas migrate hash 
.PHONY: migrate
migrate:
	$(DOCKER) run --rm --net=host \
	-v $(PWD)/internal/migrations:/migrations \
	arigaio/atlas migrate apply \
	--url "postgres://postgres:s3cr3t@localhost:5432/products?sslmode=disable"