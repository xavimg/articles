FROM golang:1.19.2-alpine

WORKDIR /app

# current directory to /app
COPY . .

RUN go build -o main main.go

EXPOSE 4007

CMD ["/app/main"]