# What is it?

httpit is a CLI tool designed to perform **Automated API** and **End-to-End (E2E) Testing** using the HTTP protocol.

> [!NOTE]
> **Project Status**: This project is currently in active development (Alpha). Core functionality for executing requests and validating responses is implemented, but features may change.

## Using Go Run
```bash
go run cmd/httpit/main.go examples/test-scheme-example.json
```

## Using Docker Image
```bash
docker run --rm -v ./examples/test-scheme-example.json:/tmp/test.json httpit /tmp/test.json
```
