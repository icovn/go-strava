######## Build API ------------------------------------------------------------
# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Nguyen Quanng Huy <quanghuy.ico@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go-strava-api/go.mod go-strava-api/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY go-strava-api .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


######## Build WEB ------------------------------------------------------------
FROM node:alpine as builder-web

WORKDIR /app

COPY go-strava-web/package*.json ./

# installs dependencies directly from package-lock.json and uses package.json
# only to validate that there are no mismatched versions
RUN npm ci

COPY go-strava-web .

RUN npm run build


######## Start a new stage from scratch ####### -------------------------------
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

RUN mkdir /root/build

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy web binary file from the previous stage
COPY --from=builder-web /app/build ./build

# Command to run the executable
CMD ["./main"]