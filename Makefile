build:
	@go build

run:
	@source .env && go run main.go

clean:
	@go clean

.PHONY: build run clean
