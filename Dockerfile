FROM golang:1.21.3-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/api/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY --from=builder /app/public /app/public
COPY --from=builder /app/scripts /app/scripts

EXPOSE 8080

CMD ["./main"] 