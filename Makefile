COMMIT := $(shell git rev-parse HEAD)

all:
	go build -ldflags "-X main.revision=$(COMMIT)" -o bin/gocodeit

clean:
	rm -rf bin

test:
	go test $(glide novendor)
