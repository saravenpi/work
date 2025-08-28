BINARY_NAME = work
GO_FILES = $(shell find . -name '*.go' -type f)

all: build

build: $(BINARY_NAME)

$(BINARY_NAME): $(GO_FILES)
	go build -o $(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

mod-tidy:
	go mod tidy

install: build
	cp $(BINARY_NAME) /usr/local/bin/

uninstall:
	rm -f /usr/local/bin/$(BINARY_NAME)

fclean: clean

re: fclean build

.PHONY: all build clean fclean test fmt vet mod-tidy install uninstall re
