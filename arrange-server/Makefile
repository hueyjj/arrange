.PHONY: all binary clean

all: run

# Run the server
run: binary
	./bin/arrange server

binary:
	mkdir -p ./bin
	go build -o ./bin/arrange ./cmd/arrange 

clean:
	rm -f ./bin/arrange