# base image needed for our application
FROM golang:1.18

# Set the environment variables
ARG DB_USER
ARG DB_PASSWORD
ENV DB_USER $DB_USER
ENV DB_PASSWORD $DB_PASSWORD

# Create and change to the app directory.
RUN mkdir -p /app

# Set the Current Working Directory inside the container
WORKDIR /app

# build the Go app
COPY . .
RUN go build -o main .

# Run the application.
CMD ["./main"]
