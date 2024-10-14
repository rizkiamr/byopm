run:
	go run main.go

build:
	go build -o byopm main.go

build-clean:
	rm -rf byopm

.PHONY: run build build-clean
