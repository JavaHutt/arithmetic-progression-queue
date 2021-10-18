tidy	:
	go mod tidy

run	:
ifdef concurrency
	go run ./cmd/main.go $(concurrency)
else
	go run ./cmd/main.go
endif
