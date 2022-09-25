FROM golang

WORKDIR /app

COPY go.mod /app

COPY go.sum /app

RUN go mod download

COPY . .

RUN go build ./cmd/main.go

EXPOSE 8080:8080

CMD ["./main"]