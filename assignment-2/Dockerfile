FROM golang:1.18 AS builder
WORKDIR /app
COPY . .

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/deploy/hacktiv-assignment-2 .
EXPOSE 8080

ENV APP_NAME=hacktiv-assignment-2
CMD ["./hacktiv-assignment-2"]
