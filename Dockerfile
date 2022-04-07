FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o sepatuku-api

EXPOSE 8080

CMD ./sepatuku-api