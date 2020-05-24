# Largely inspired by https://danishpraka.sh/2019/12/07/using-makefiles-for-go.html
GO_ARGS=-race
GO111MODULES=on

.PHONY: build
build:
## build: build this application
	go build  ${GO_ARGS} -o znaczek github.com/jpalczewski/znaczek/cmd/znaczek 

.PHONY: clean
clean:
## clean: clean the application
	go clean

.PHONY: setup
setup:
## setup: setup go modules
	@go mod tidy \
		&& go mod vendor

.PHONY: test
test:
## test: run tests
	go test -v -count=1 -race ./...


.PHONY: install
install:
## install: install the binary
	go install


.PHONY: test-release
test-release:
	goreleaser --snapshot --skip-publish --rm-dist

.PHONY:
	goreleaser

.PHONY: help
## help: prints this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
