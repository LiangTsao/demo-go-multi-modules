MODULES = moduleA moduleB

.PHONY: all build build-modules build-app test clean run help

all: build

build: build-modules build-app

build-modules:
	for module in $(MODULES); do \
		(cd $$module && go build ./...); \
	done

build-app:
	mkdir -p bin
	go build -o bin/app1 ./cmd/app1

run: build-app
	./bin/app1

test:
	for module in $(MODULES); do \
		(cd $$module && go test ./...); \
	done

clean:
	for module in $(MODULES); do \
		(cd $$module && go clean); \
	done
	rm -rf bin

help:
	@echo "Available targets:"
	@echo "  all     - Build all components (default target)"
	@echo "  build   - Build modules and application"
	@echo "  run     - Build and run the application"
	@echo "  test    - Run tests in all modules"
	@echo "  clean   - Clean build artifacts"
	@echo "  help    - Display this help message"