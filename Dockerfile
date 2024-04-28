FROM golang:1.17-alpine3.14

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o backend_task .

CMD [ "./backend_task" ]

EXPOSE 8080
