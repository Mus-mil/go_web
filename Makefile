
all: clear build
	./yoldiz

run:
	go run cmd/app/main.go

build:
	go build -o yoldiz cmd/app/main.go

clear:
	rm -rf yoldiz