# Create multistage docker container image 
# create docker container from base image 

FROM golang:1.22.5-alpine as builder

RUN mkdir /app

WORKDIR /app

RUN . /app

RUN CGO_ENABLED=0 go build -o authAPI ./cmd/api

RUN chmod +x /app/authAPI


# Create tiny docker container image from builder 
FROM alpine:latest 

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/authAPI /app
CMD [ "/app/authAPI" ]

