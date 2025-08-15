FROM golang:tip-alpine3.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build cmd/forecast/main.go

EXPOSE 8080

CMD ["./main"]