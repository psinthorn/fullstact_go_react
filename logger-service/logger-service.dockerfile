# multi stage dockerfile for logger-service
# stage 1: build the go application
FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .  

RUN CGO_ENABLED=0 go build -o loggerAPI ./cmd/api

# stage 2: build the final image
FROM alpine:3.13

WORKDIR /app

COPY --from=build /app/loggerAPI .

CMD ["./loggerAPI"]



