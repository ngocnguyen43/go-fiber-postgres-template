# Use a smaller base image
FROM alpine:3.18

# Install curl and ca-certificates
RUN apk add --no-cache curl ca-certificates

# Install Atlasgo using curl
RUN curl -sSf https://atlasgo.sh | sh

# Copy the migrations folder into the container
COPY ../migrations /migrations

# Set up the entry point for the container
CMD ["sh", "-c", "atlas migrate apply --url postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:5432/$DB_DATABASE?sslmode=disable"]
