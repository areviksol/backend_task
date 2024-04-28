FROM golang:1.17 as builder

WORKDIR /app

COPY . .

RUN go mod download && go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8000

CMD ["./main"]
