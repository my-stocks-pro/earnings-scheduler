
#GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main -ldflags '-w -s' main.go

BINARY=earnings-scheduler

GOOS=linux
GOARCH=amd64
CGO_ENABLED=0

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFLAGS=-ldflags '-w -s'
DOCKER=docker
DOCKERBUILD=$(DOCKER) build
DOCKERRUN=$(DOCKER) run

GODEP=dep
NEWDEP=$(GODEP) ensure

all: test go-build docker-build

test:
	$(GOTEST) -v ./...

go-build:
	$(NEWDEP)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) $(GOBUILD) -o $(BINARY) $(GOFLAGS) main.go

docker-build:
	$(DOCKERBUILD) --no-cache -t $(BINARY_NAME) .

run:
	$(DOCKERRUN) \
	--rm \
	-d \
	-p 8002:8002 \
	--name=$(BINARY) \
	$(BINARY)

clean:
	$(GOCLEAN)
	rm -f $(BINARY)