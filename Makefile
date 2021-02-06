BINDIR := $(CURDIR)/bin
BINNAME := prometheus-custom-metrics-example

SRC        := $(shell find . -type f -name '*.go' -print)

.PHONY: build
build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME): $(SRC)
	go build -o $(BINDIR)/$(BINNAME) ./cmd/example/

.PHONY: clean
clean:
	rm -rf bin/
