FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

COPY ./entrypoint.sh ./

RUN chmod +x ./entrypoint.sh

CMD ["/bin/sh", "-c", "/app/entrypoint.sh"]

