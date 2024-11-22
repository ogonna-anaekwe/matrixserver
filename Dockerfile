FROM golang:1.21.1

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o /go/bin/main /app/cmd/main.go

EXPOSE 8080

CMD ["/go/bin/main"]