FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go build -o main .

EXPOSE 8000

CMD "air"