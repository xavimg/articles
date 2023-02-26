FROM golang:1.19.2-alpine

WORKDIR /app

COPY . .

RUN go build -o main main.go

EXPOSE 4007

CMD ["/app/main"]