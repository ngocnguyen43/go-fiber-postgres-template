FROM golang:1.22.6-alpine AS build

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init  -d ./cmd/api,./internal --parseDependency --parseInternal --parseDepth=10

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o run cmd/api/main.go

# Moving the binary to the 'final Image' to make it smaller
FROM alpine

WORKDIR /app

COPY --from=build /build/run .

CMD ["/app/run"]