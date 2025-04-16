MODULES = moduleA moduleB

.PHONY: all build test clean

all: build

build:
	for module in $(MODULES); do \
		(cd $$module && go build ./...); \
	done

test:
	for module in $(MODULES); do \
		(cd $$module && go test ./...); \
	done

clean:
	for module in $(MODULES); do \
		(cd $$module && go clean); \
	done