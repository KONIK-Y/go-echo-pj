FROM golang:1.23.2-alpine3.20 AS build

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

RUN apk update && \
    apk add --no-cache curl gnupg && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz -o migrate.tar.gz && \
    tar -xvzf migrate.tar.gz && \
    mv migrate /usr/local/bin/migrate && \
    chmod +x /usr/local/bin/migrate && \
    rm -rf migrate.tar.gz /var/cache/apk/*
    
COPY src/ /app/src

RUN go build -o main ./src

EXPOSE 8080

CMD ["./main"]