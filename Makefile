run:
	go run cmd/byopm/main.go

build:
	go build -o app cmd/byopm/main.go

build-clean:
	rm -rf app

.PHONY: run build build-clean
