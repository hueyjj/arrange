.PHONY: all binary clean

all: binary

binary:
	mkdir -p ./bin
	go build -o ./bin/arrange ./cmd/arrange 

clean:
	rm -f ./bin/arrange