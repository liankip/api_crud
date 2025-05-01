FROM golang:1.24-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/google/wire/cmd/wire@latest

CMD ["air", "-c", ".air.toml"]
