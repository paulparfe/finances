FROM golang:1.23.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/app/main.go

# установка и перенос goose
RUN go install -tags 'postgres' github.com/pressly/goose/v3/cmd/goose@latest
#RUN cp /go/bin/goose /app/goose


FROM ubuntu:22.04

WORKDIR /app

RUN apt update && apt install -y ca-certificates
# RUN apt update && apt install -y mc net-tools htop

COPY --from=builder /app/main /app/main
COPY --from=builder /app/.env.example /app/.env
COPY --from=builder /app/migrations /app/migrations

# переносим goose
COPY --from=builder /go/bin/goose /usr/local/bin/goose
RUN chmod +x /usr/local/bin/goose

EXPOSE 8080

# запуск
CMD ["/app/main"]
