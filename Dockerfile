FROM golang:1.22.4-alpine AS build

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o run cmd/api/main.go

# Moving the binary to the 'final Image' to make it smaller
FROM alpine

WORKDIR /app

COPY --from=build /build/run .

ENV PORT 8000
ENV DB_HOST localhost
ENV DB_PORT 5432
ENV DB_DATABASE webhook
ENV DB_USERNAME minhngocnguyen
ENV DB_PASSWORD minhngoc.403
ENV DB_SCHEMA public

CMD ["/app/run"]