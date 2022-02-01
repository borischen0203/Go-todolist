FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/app main.go

EXPOSE 8000

CMD ["./bin/app"]
