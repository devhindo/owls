# Base image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the source code to the working directory
COPY . .

# Build the application
RUN go build -o main .

# Expose the port on which the server will run
EXPOSE 8080

# Set the command to run the server
CMD ["./main"]
