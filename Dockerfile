# --- Build Stage ---
# Start with the Go image that contains the compiler and tools.
# We give this stage a name, "builder", so we can reference it later.
FROM golang:1.24-alpine AS builder

# Set the working directory for the build.
WORKDIR /app

# Copy the module files first. This is a caching optimization.
# If these files don't change, Docker won't re-download dependencies.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the source code.
COPY . .

# Build the application, creating a statically linked binary.
# This ensures it can run on its own without system dependencies.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# --- Run Stage ---
# Start from a new, extremely small base image.
# 'alpine' is a minimal Linux distribution, perfect for running single binaries.
FROM alpine:latest

# Set the working directory for the final container.
WORKDIR /root/

# Copy *only* the compiled binary from the "builder" stage.
# This is the magic of multi-stage builds.
COPY --from=builder /app/main .

# Expose port 8080 for incoming traffic.
EXPOSE 8080

# The command to run the application when the container starts.
CMD ["./main"]