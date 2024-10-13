.PHONY: test

test:
	@go test ./...

run:
	@go run main/main.go

commit: add
	@git commit

add: test
	@git add .

