SRCS := $(shell find . -name '*.go')
LINTERS := \
	golang.org/x/lint/golint \
	honnef.co/go/tools/cmd/staticcheck
APP_NAME := daff

CONFIG_FILE := config.yaml

BANNER:=\
    "\n"\
		"/**\n"\
    " * @project       $(APP_NAME)\n"\
    " */\n"\
    "\n"

## build.linux			: Build application for Linux runtime
.PHONY: build.linux
build.linux:
	env GOOS=linux go build -ldflags="-s -w" -o bin/$(APP_NAME) main.go

## build.mac			: Build application for Mac runtime
.PHONY: build.mac
build.mac:
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/$(APP_NAME) main.go

## clean				: Clean application objects
.PHONY: clean
clean:
	go clean -i ./...

## deps				: Download dependencies
.PHONY: deps
deps:
	go get -d -v ./...

## golint				: Run golint on all *.go files
.PHONY: golint
golint: 
	for file in $(SRCS); do \
		golint $${file}; \
		if [ -n "$$(golint $${file})" ]; then \
			exit 1; \
		fi; \
	done

## help				: Show all available make targets
.PHONY : help
help : 
	@echo $(BANNER)
	@echo \ Make targets
	@echo -----------------------------
	@sed -n 's/^##//p' Makefile | sort

## install			: Install dependencies
.PHONY: install
install: deps
	go install ./...

## lint				: Run golint, vet and staticcheck
.PHONY: lint
lint: golint vet staticcheck

## run				: Run application from main.go
.PHONY: run
run:
	go run main.go

## start				: Start application
.PHONY: start
start:
	./bin/$(APP_NAME) -f $(CONFIG_FILE) -t $(BOT_TOKEN)

## staticcheck			: Run staticcheck
.PHONY: staticcheck
staticcheck:
	staticcheck ./...

## test				: Run tests
.PHONY: test
test:
	go test -v -race -coverprofile=coverage.out ./...

## test.coverage			: Show test coverage report
.PHONY: test.coverage
test.coverage:
	go tool cover -html=coverage.out

## testdeps			: Download test dependencies
.PHONY: testdeps
testdeps:
	go get -d -v -t ./...
	go get -v $(LINTERS)

## vet				: Run vet
.PHONY: vet
vet:
	go vet ./...
