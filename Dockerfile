FROM golang:latest

LABEL maintainer="Lakshitha Herath"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o main .

EXPOSE 8080

# Command to run the executable
CMD ["./main"]