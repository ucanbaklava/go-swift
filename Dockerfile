FROM golang:1.21.6

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

ENV CGO_ENABLED = 1

COPY . .

RUN go build -o myapp

CMD ["./myapp"]