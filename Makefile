run:
	go run cmd/byopm/main.go

build:
	go build -o build/byopm cmd/byopm/main.go

build-clean:
	rm -rf build/byopm

.PHONY: run build build-clean
