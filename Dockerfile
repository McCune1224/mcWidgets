# Build the Go application
FROM golang:1.23 AS backend

WORKDIR /app

# Copy the Tailwind CSS files
COPY static/css/styles.css ./static/css/
COPY tailwind.config.js ./
#

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o myapp .

# Stage 3: Create the final image
FROM alpine:latest

WORKDIR /app

# Copy the built Go binary and static files from previous stages
COPY --from=backend /app/myapp .
COPY --from=frontend /app/static ./static

# Install the necessary packages for running Go applications
RUN apk add --no-cache ca-certificates

# Expose the application port
EXPOSE 8080

# Run the Go application
CMD ["./myapp"]
