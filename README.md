# What is it?

httpit is an application designed to perform integration testing using the http protocol.

## Using Go Run
```bash
go run cmd/httpit.go test-scheme-example.json
```

## Using Docker Image
```bash
docker run --rm -v ./test-scheme-example.json:/tmp/test.json apagar /tmp/test.json
```
