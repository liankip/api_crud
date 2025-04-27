FROM golang:latest
RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go get github.com/gofiber/fiber/v2
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/mysql

COPY go.mod ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]