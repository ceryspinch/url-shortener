# Use an official Go runtime as a parent image
FROM golang:1.16

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build -o url-shortener .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./url-shortener"]