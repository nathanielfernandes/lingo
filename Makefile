NAME=lingo
BUILD=build/

run:
	@go run main.go

.PHONY: build
build:
	@mkdir -p $(BUILD)
	@go build -v -o $(BUILD)/$(NAME) main.go
