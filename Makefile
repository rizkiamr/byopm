run:
	go run cmd/byopm/main.go

build:
	go build -o byopm cmd/byopm/main.go

build-clean:
	rm -rf byopm

.PHONY: run build build-clean
