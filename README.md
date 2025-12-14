# What is it?

httpit is a CLI tool designed to perform **Automated API** and **End-to-End (E2E) Testing** using the HTTP protocol.

## Using Go Run
```bash
go run cmd/httpit/main.go examples/test-scheme-example.json
```

## Using Docker Image
```bash
docker run --rm -v ./examples/test-scheme-example.json:/tmp/test.json httpit /tmp/test.json
```
