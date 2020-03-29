BINARY_NAME=tictacgo

all: help
help: Makefile
	@echo "Makefile \nChoose a command to run:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
## build: builds the app and generates a binary called tictacgo
build: 
	go build -o $(BINARY_NAME)
## test: runs tests
test: build 
	go test
## run: runs the app
run: build 
	./$(BINARY_NAME)