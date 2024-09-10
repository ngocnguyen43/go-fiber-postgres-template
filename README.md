# Project go-fiber-postgres-template

One Paragraph of project description goes here

## Getting Started


## Generate Swagger Docs (require swaggo [link](https://github.com/swaggo/swag))

```bash
swag init -d ./cmd/api,./ --parseDependency --parseInternal
```

## Migration (require atlasgo [link](https://atlasgo.io/))

### Generate sql file
```bash
atlas migrate diff  <migration_name>  --env gorm 
```
**_NOTE:_** _Add models struct inside `cmd/loader/main.go`_

### Apply migartions
```bash
atlas migrate apply --url "postgres://localhost:postgres@:5432/go-fiber?search_path=public&sslmode=disable"
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
docker compose up
```bash
make up
```
docker compose down
```bash
make down
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```