# INFO: Project variables
.PHONY: run start build commit add test getcover cover clean

default: help

# INFO: Project commands
run:
	@go run main/main.go

start:
	@./build/arsg

startclean: build
	@./build/arsg

build: test clean
	@mkdir -p build
	@go build -buildmode=exe -o ./build/arsg -trimpath ./main/main.go
	@upx --best --lzma ./build/arsg

install: build
	@chmod a+x ./scripts/install.sh
	@./scripts/install.sh

uninstall:
	@chmod a+x ./scripts/install.sh
	@./scripts/uninstall.sh

commit: add
	@git commit

add: test
	@git add .

test:
	@go test ./...

clean:
	@rm -rf ./build/

getcover: cover
	@go tool cover -html=./out/coverage.out -o ./out/coverage.html

cover:
	@mkdir -p out
	@go test -coverprofile=./out/coverage.out ./...

help:
	@echo "run       run the app"
	@echo "start     build and run the app"
	@echo "build     build the app"
	@echo "commit    commit the changes"
	@echo "add       add the changes"
	@echo "test      run tests"
	@echo "clean     clean the build directory"
	@echo "getcover  get the coverage report"
	@echo "cover     run the tests and generate the coverage report"
