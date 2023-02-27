# Build stage
FROM golang:1.19.2-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o main main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main .
COPY --from=build /app/config/config.yml ./config/

EXPOSE 4007

CMD ["./main"]
