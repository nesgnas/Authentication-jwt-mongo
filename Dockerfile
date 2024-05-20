# syntax=docker/dockerfile:1

FROM golang:1.22

# Set destination for COPY
WORKDIR /jwt

COPY . .

RUN go mod download

RUN go build -o bin .

ENTRYPOINT ["/jwt/bin"]

