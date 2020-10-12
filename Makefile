
GIT_COMMIT := $(shell git rev-list -1 HEAD)
GOLDFLAGS += -X main.GitCommit=$(GIT_COMMIT)
GOFLAGS = -ldflags "$(GOLDFLAGS)" -trimpath

default: build

build:
	go build $(GOFLAGS) -o kamel pkg/cmd/*.go
clean:
	go clean
	rm kamel	
