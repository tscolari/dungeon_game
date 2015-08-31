all: test build

build:
	go build ./main.go

test:
	ginkgo -r
