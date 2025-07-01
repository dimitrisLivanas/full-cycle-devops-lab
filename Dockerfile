# Stage 1: Start with a base image that has the Go language installed.
# 'golang:1.22-alpine' is an official image that's small and efficient.
FROM golang:1.24-alpine

# Set the working directory inside the container.
# This is where our code will live.
WORKDIR /app

# Copy all the local source code into the container's working directory.
COPY . .

# Compile the Go application inside the container.
# -o main specifies that the output executable should be named 'main'.
RUN go build -o main .

# This instruction tells Docker that the container will listen on port 8080 at runtime.
# This is for documentation purposes.
EXPOSE 8080

# This is the command that will be run when the container starts.
# It executes our compiled program.
CMD ["./main"]