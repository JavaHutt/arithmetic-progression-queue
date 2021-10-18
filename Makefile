tidy	:
	go mod tidy

docker	:
	docker build -t arithmetic-progression-app .

run	:
ifdef concurrency
	go run ./cmd/main.go $(concurrency)
else
	go run ./cmd/main.go
endif
