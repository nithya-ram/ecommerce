FROM golang:1.25.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ecommerce ./main.go


FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/ecommerce .

COPY view ./view


EXPOSE 8080

CMD ["./ecommerce"]