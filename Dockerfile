# Start from golang base image
FROM golang:1.20-alpine as builder

# Add Maintainer info
LABEL maintainer="Alam <nc.alamsyah@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /github.com/ncalamsyah/e-commerce

# Download necessary Go modules
COPY . .
EXPOSE 9200
RUN go mod tidy
RUN go build
CMD go run main.go