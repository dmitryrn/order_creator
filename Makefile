ifeq ($(CICD_MODE),true)
	ENVFILE_BASE=.env.cicd
else
	ENVFILE_BASE=.env.example
endif

.env:
	cp $(ENVFILE_BASE) .env

.PHONY: deps
deps: .env
	docker compose run --rm go-custom go mod tidy

.PHONY: test
test: .env
	docker compose run --rm go-custom go test -v ./...

.PHONY: build
build: .env
	docker compose run --rm go-custom make _build

.PHONY: _build
_build:
	rm -rf bin
	@for dir in $(wildcard functions/*/) ; do \
		fxn=$$(basename $$dir) ; \
		GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o bin/$$fxn/bootstrap functions/$$fxn/*.go ; \
		zip -j bin/$$fxn.zip bin/$$fxn/bootstrap ; \
	done

.PHONY: deploy
deploy: .env bin
	docker compose run --rm serverless sls deploy

.PHONY: fmt
fmt: .env
	docker compose run --rm go-custom go fmt ./...

.PHONY: mocks
mocks: .env
	docker compose run --rm mockery
