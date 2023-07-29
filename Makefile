build:
	@go build -o ./bin/sstatus

run: build
	./bin/sstatus