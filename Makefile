
all: clear build
	./yoldiz

run:
	go run cmd/main.go

build:
	go build -o yoldiz cmd/main.go

clear:
	rm -rf yoldiz