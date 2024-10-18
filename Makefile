CC = go
BIN_NAME = httpit

build:
	$(CC) build -o $(BIN_NAME) cmd/httpit.go
clean:
	rm -f $(BIN_NAME)
