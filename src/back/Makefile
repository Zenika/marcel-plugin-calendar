IMAGE_NAME = marcel_agenda
VERSION = 1.0

GOCMD = go
GOCLEAN = $(GOCMD) clean
GOFMT = $(GOCMD) fmt
GOTEST = $(GOCMD) test
GOBUILD = $(GOCMD) build -o $(IMAGE_NAME) .
SAVE = docker save $(IMAGE_NAME):$(VERSION) > ../../dist/back/$(IMAGE_NAME)_$(VERSION).tar

all: test
	$(GOCLEAN)
	$(GOFMT)
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(IMAGE_NAME) .
	docker build -t $(IMAGE_NAME):$(VERSION) .
	$(SAVE)

test:
	$(GOTEST) -v -covermode=count ./...

build: test
	$(GOBUILD)

clean:
	$(GOCLEAN)

quick: test
	$(GOCLEAN)
	$(GOFMT)
	$(GOBUILD)

q: quick

push:
	docker push zenika/$(IMAGE_NAME):$(VERSION)

save:
	$(SAVE)
