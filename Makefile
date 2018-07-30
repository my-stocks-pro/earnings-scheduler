
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

all: go-build docker-build clean

go-build:
	@echo "Golang build executable"
	$(NEWDEP)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) $(GOBUILD) -o $(BINARY) $(GOFLAGS) main.go

docker-build:
	@echo "Docker build service"
	$(DOCKERBUILD) --no-cache -t $(BINARY) .

run:
	@echo "Docker run service"
	$(DOCKERRUN) \
	-q \
	--rm \
	-dt \
	-p 8002:8002 \
	--name=$(BINARY) \
	$(BINARY)

clean:
	@echo "Clean"
	$(GOCLEAN)
	rm -f $(BINARY)