FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build ./cmd/main.go

CMD ["./main"]