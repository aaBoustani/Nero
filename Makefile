build:
	go build -o cmd/briq

run: build
	./cmd/briq
