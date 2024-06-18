# Project go-fiber-postgres-template

One Paragraph of project description goes here

## Getting Started


## Generate Swagger Docs

default API docs<br>
http://localhost:8080/docs

generate APIs docs
```bash
swag init -d ./cmd/api,./
```

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```